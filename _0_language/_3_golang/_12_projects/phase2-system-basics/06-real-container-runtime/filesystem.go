package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

// FilesystemManager manages container filesystems
type FilesystemManager struct {
	basePath    string
	containerID string
}

// ImageLayer represents a filesystem layer
type ImageLayer struct {
	ID       string `json:"id"`
	Parent   string `json:"parent"`
	Path     string `json:"path"`
	Size     int64  `json:"size"`
	Created  string `json:"created"`
	Digest   string `json:"digest"`
}

// ContainerFilesystem represents a container's filesystem
type ContainerFilesystem struct {
	ID          string       `json:"id"`
	Image       string       `json:"image"`
	Layers      []ImageLayer `json:"layers"`
	MergedPath  string       `json:"merged_path"`
	WorkPath    string       `json:"work_path"`
	UpperPath   string       `json:"upper_path"`
	LowerPath   string       `json:"lower_path"`
}

// NewFilesystemManager creates a new filesystem manager
func NewFilesystemManager(containerID string) *FilesystemManager {
	return &FilesystemManager{
		basePath:    "/var/lib/containers",
		containerID: containerID,
	}
}

// CreateContainerFilesystem creates a container filesystem
func (fm *FilesystemManager) CreateContainerFilesystem(image string) (*ContainerFilesystem, error) {
	fmt.Printf("💾 Creating container filesystem for image: %s\n", image)
	
	// Create base directories
	if err := fm.createBaseDirectories(); err != nil {
		return nil, fmt.Errorf("failed to create base directories: %w", err)
	}
	
	// Create container filesystem structure
	containerFS := &ContainerFilesystem{
		ID:         fm.containerID,
		Image:      image,
		MergedPath: filepath.Join(fm.basePath, "merged", fm.containerID),
		WorkPath:   filepath.Join(fm.basePath, "work", fm.containerID),
		UpperPath:  filepath.Join(fm.basePath, "upper", fm.containerID),
		LowerPath:  filepath.Join(fm.basePath, "lower", fm.containerID),
	}
	
	// Create image layers
	if err := fm.createImageLayers(containerFS, image); err != nil {
		return nil, fmt.Errorf("failed to create image layers: %w", err)
	}
	
	// Create overlay filesystem
	if err := fm.createOverlayFilesystem(containerFS); err != nil {
		return nil, fmt.Errorf("failed to create overlay filesystem: %w", err)
	}
	
	fmt.Printf("✅ Container filesystem created at: %s\n", containerFS.MergedPath)
	return containerFS, nil
}

// createBaseDirectories creates the base directory structure
func (fm *FilesystemManager) createBaseDirectories() error {
	dirs := []string{
		"images", "layers", "merged", "work", "upper", "lower", "tmp",
	}
	
	for _, dir := range dirs {
		path := filepath.Join(fm.basePath, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
	}
	
	return nil
}

// createImageLayers creates image layers
func (fm *FilesystemManager) createImageLayers(containerFS *ContainerFilesystem, image string) error {
	fmt.Println("  Creating image layers...")
	
	// For demo purposes, create a basic filesystem structure
	// In a real implementation, this would extract image layers from a registry
	
	// Create base layer
	baseLayer := ImageLayer{
		ID:      "base-" + fm.containerID,
		Parent:  "",
		Path:    filepath.Join(fm.basePath, "layers", "base"),
		Size:    0,
		Created: "2024-01-01T00:00:00Z",
		Digest:  "sha256:base",
	}
	
	if err := fm.createBaseLayer(&baseLayer); err != nil {
		return fmt.Errorf("failed to create base layer: %w", err)
	}
	
	// Create image layer
	imageLayer := ImageLayer{
		ID:      "image-" + fm.containerID,
		Parent:  baseLayer.ID,
		Path:    filepath.Join(fm.basePath, "layers", fm.containerID),
		Size:    0,
		Created: "2024-01-01T00:00:00Z",
		Digest:  "sha256:" + fm.containerID,
	}
	
	if err := fm.createImageLayer(&imageLayer, image); err != nil {
		return fmt.Errorf("failed to create image layer: %w", err)
	}
	
	containerFS.Layers = []ImageLayer{baseLayer, imageLayer}
	containerFS.LowerPath = baseLayer.Path
	
	return nil
}

// createBaseLayer creates the base layer
func (fm *FilesystemManager) createBaseLayer(layer *ImageLayer) error {
	fmt.Printf("    Creating base layer: %s\n", layer.ID)
	
	// Create layer directory
	if err := os.MkdirAll(layer.Path, 0755); err != nil {
		return fmt.Errorf("failed to create layer directory: %w", err)
	}
	
	// Create basic filesystem structure
	dirs := []string{
		"bin", "sbin", "usr", "etc", "var", "tmp", "proc", "sys", "dev",
		"usr/bin", "usr/sbin", "usr/lib", "usr/lib64",
		"var/log", "var/lib", "var/tmp",
		"etc/init.d", "etc/rc.d", "etc/systemd",
	}
	
	for _, dir := range dirs {
		path := filepath.Join(layer.Path, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
	}
	
	// Create essential files
	files := map[string]string{
		"etc/hostname": "container",
		"etc/hosts":    "127.0.0.1 localhost\n::1 localhost",
		"etc/passwd":   "root:x:0:0:root:/root:/bin/bash\nnobody:x:65534:65534:nobody:/nonexistent:/usr/sbin/nologin",
		"etc/group":    "root:x:0:\nnobody:x:65534:",
		"etc/fstab":    "proc /proc proc defaults 0 0\ntmpfs /tmp tmpfs defaults 0 0",
	}
	
	for file, content := range files {
		path := filepath.Join(layer.Path, file)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %w", path, err)
		}
	}
	
	return nil
}

// createImageLayer creates an image-specific layer
func (fm *FilesystemManager) createImageLayer(layer *ImageLayer, image string) error {
	fmt.Printf("    Creating image layer: %s for %s\n", layer.ID, image)
	
	// Create layer directory
	if err := os.MkdirAll(layer.Path, 0755); err != nil {
		return fmt.Errorf("failed to create layer directory: %w", err)
	}
	
	// Create image-specific files based on image type
	if strings.Contains(image, "ubuntu") || strings.Contains(image, "debian") {
		return fm.createUbuntuLayer(layer)
	} else if strings.Contains(image, "alpine") {
		return fm.createAlpineLayer(layer)
	} else if strings.Contains(image, "centos") || strings.Contains(image, "rhel") {
		return fm.createCentOSLayer(layer)
	} else {
		return fm.createGenericLayer(layer)
	}
}

// createUbuntuLayer creates an Ubuntu-based layer
func (fm *FilesystemManager) createUbuntuLayer(layer *ImageLayer) error {
	fmt.Println("      Creating Ubuntu layer...")
	
	// Create Ubuntu-specific directories
	dirs := []string{
		"usr/share", "usr/local", "opt", "home", "root",
		"usr/share/doc", "usr/share/man", "usr/share/info",
		"var/cache", "var/spool", "var/mail",
	}
	
	for _, dir := range dirs {
		path := filepath.Join(layer.Path, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
	}
	
	// Create Ubuntu-specific files
	files := map[string]string{
		"etc/os-release": `PRETTY_NAME="Ubuntu 22.04 LTS"
NAME="Ubuntu"
VERSION_ID="22.04"
VERSION="22.04 LTS"
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
VERSION_CODENAME=jammy
UBUNTU_CODENAME=jammy`,
		"etc/debian_version": "22.04",
		"etc/issue":          "Ubuntu 22.04 LTS \\n \\l",
	}
	
	for file, content := range files {
		path := filepath.Join(layer.Path, file)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %w", path, err)
		}
	}
	
	return nil
}

// createAlpineLayer creates an Alpine-based layer
func (fm *FilesystemManager) createAlpineLayer(layer *ImageLayer) error {
	fmt.Println("      Creating Alpine layer...")
	
	// Create Alpine-specific directories
	dirs := []string{
		"usr/share", "usr/local", "opt", "home", "root",
		"usr/share/apk", "var/cache/apk",
	}
	
	for _, dir := range dirs {
		path := filepath.Join(layer.Path, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
	}
	
	// Create Alpine-specific files
	files := map[string]string{
		"etc/os-release": `NAME="Alpine Linux"
ID=alpine
VERSION_ID=3.18.0
PRETTY_NAME="Alpine Linux v3.18"
HOME_URL="https://alpinelinux.org/"
BUG_REPORT_URL="https://gitlab.alpinelinux.org/alpine/aports/-/issues"`,
		"etc/alpine-release": "3.18.0",
	}
	
	for file, content := range files {
		path := filepath.Join(layer.Path, file)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %w", path, err)
		}
	}
	
	return nil
}

// createCentOSLayer creates a CentOS-based layer
func (fm *FilesystemManager) createCentOSLayer(layer *ImageLayer) error {
	fmt.Println("      Creating CentOS layer...")
	
	// Create CentOS-specific directories
	dirs := []string{
		"usr/share", "usr/local", "opt", "home", "root",
		"usr/share/doc", "usr/share/man", "usr/share/info",
		"var/cache", "var/spool", "var/mail", "var/lib/rpm",
	}
	
	for _, dir := range dirs {
		path := filepath.Join(layer.Path, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
	}
	
	// Create CentOS-specific files
	files := map[string]string{
		"etc/os-release": `NAME="CentOS Linux"
VERSION="8"
ID="centos"
ID_LIKE="rhel fedora"
VERSION_ID="8"
PLATFORM_ID="platform:el8"
PRETTY_NAME="CentOS Linux 8"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:centos:centos:8"
HOME_URL="https://www.centos.org/"
BUG_REPORT_URL="https://bugs.centos.org/"
CENTOS_MANTISBT_PROJECT="CentOS-8"
CENTOS_MANTISBT_PROJECT_VERSION="8"`,
		"etc/redhat-release": "CentOS Linux release 8.0.1905 (Core)",
	}
	
	for file, content := range files {
		path := filepath.Join(layer.Path, file)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %w", err)
		}
	}
	
	return nil
}

// createGenericLayer creates a generic layer
func (fm *FilesystemManager) createGenericLayer(layer *ImageLayer) error {
	fmt.Println("      Creating generic layer...")
	
	// Create generic directories
	dirs := []string{
		"usr/share", "usr/local", "opt", "home", "root",
	}
	
	for _, dir := range dirs {
		path := filepath.Join(layer.Path, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", path, err)
		}
	}
	
	return nil
}

// createOverlayFilesystem creates an overlay filesystem
func (fm *FilesystemManager) createOverlayFilesystem(containerFS *ContainerFilesystem) error {
	fmt.Println("  Creating overlay filesystem...")
	
	// Create overlay directories
	dirs := []string{containerFS.MergedPath, containerFS.WorkPath, containerFS.UpperPath}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}
	
	// Mount overlay filesystem
	if err := fm.mountOverlay(containerFS); err != nil {
		return fmt.Errorf("failed to mount overlay: %w", err)
	}
	
	fmt.Printf("    Overlay filesystem mounted at: %s\n", containerFS.MergedPath)
	return nil
}

// mountOverlay mounts the overlay filesystem
func (fm *FilesystemManager) mountOverlay(containerFS *ContainerFilesystem) error {
	// Create overlay mount options
	options := fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s",
		containerFS.LowerPath,
		containerFS.UpperPath,
		containerFS.WorkPath)
	
	// Mount overlay filesystem
	if err := syscall.Mount("overlay", containerFS.MergedPath, "overlay", 0, options); err != nil {
		return fmt.Errorf("overlay mount failed: %w", err)
	}
	
	return nil
}

// MountVolume mounts a volume in the container
func (fm *FilesystemManager) MountVolume(containerFS *ContainerFilesystem, hostPath, containerPath string) error {
	fmt.Printf("💾 Mounting volume: %s -> %s\n", hostPath, containerPath)
	
	// Create container path
	fullContainerPath := filepath.Join(containerFS.MergedPath, containerPath)
	if err := os.MkdirAll(fullContainerPath, 0755); err != nil {
		return fmt.Errorf("failed to create container path: %w", err)
	}
	
	// Mount the volume
	if err := syscall.Mount(hostPath, fullContainerPath, "", syscall.MS_BIND, ""); err != nil {
		return fmt.Errorf("volume mount failed: %w", err)
	}
	
	fmt.Printf("✅ Volume mounted successfully\n")
	return nil
}

// UnmountVolume unmounts a volume
func (fm *FilesystemManager) UnmountVolume(containerFS *ContainerFilesystem, containerPath string) error {
	fullContainerPath := filepath.Join(containerFS.MergedPath, containerPath)
	
	// Unmount the volume
	if err := syscall.Unmount(fullContainerPath, 0); err != nil {
		return fmt.Errorf("volume unmount failed: %w", err)
	}
	
	return nil
}

// CleanupFilesystem cleans up the container filesystem
func (fm *FilesystemManager) CleanupFilesystem(containerFS *ContainerFilesystem) error {
	fmt.Printf("🧹 Cleaning up filesystem for container: %s\n", fm.containerID)
	
	// Unmount overlay filesystem
	if err := syscall.Unmount(containerFS.MergedPath, 0); err != nil {
		fmt.Printf("Warning: failed to unmount overlay: %v\n", err)
	}
	
	// Remove container directories
	dirs := []string{
		containerFS.MergedPath,
		containerFS.WorkPath,
		containerFS.UpperPath,
	}
	
	for _, dir := range dirs {
		if err := os.RemoveAll(dir); err != nil {
			fmt.Printf("Warning: failed to remove directory %s: %v\n", dir, err)
		}
	}
	
	return nil
}

// GetFilesystemInfo returns filesystem information
func (fm *FilesystemManager) GetFilesystemInfo(containerFS *ContainerFilesystem) (*FilesystemInfo, error) {
	info := &FilesystemInfo{
		ContainerID: containerFS.ID,
		Image:       containerFS.Image,
		MergedPath:  containerFS.MergedPath,
		Layers:      len(containerFS.Layers),
	}
	
	// Get filesystem usage
	if stat, err := os.Stat(containerFS.MergedPath); err == nil {
		info.Size = stat.Size()
	}
	
	// Get layer information
	for _, layer := range containerFS.Layers {
		if stat, err := os.Stat(layer.Path); err == nil {
			info.TotalSize += stat.Size()
		}
	}
	
	return info, nil
}

// CheckOverlaySupport checks if overlay filesystem is supported
func (fm *FilesystemManager) CheckOverlaySupport() error {
	fmt.Println("🔍 Checking overlay filesystem support...")
	
	// Check if overlay module is loaded
	if data, err := os.ReadFile("/proc/filesystems"); err == nil {
		if !strings.Contains(string(data), "overlay") {
			return fmt.Errorf("overlay filesystem not supported")
		}
	}
	
	// Check if overlay is available
	if _, err := os.Stat("/sys/module/overlay"); err != nil {
		return fmt.Errorf("overlay module not loaded")
	}
	
	fmt.Println("✅ Overlay filesystem support available")
	return nil
}

// PrintFilesystemInfo prints filesystem information
func (fm *FilesystemManager) PrintFilesystemInfo(containerFS *ContainerFilesystem) {
	fmt.Println("💾 Filesystem Information")
	fmt.Println("========================")
	
	// Check overlay support
	if err := fm.CheckOverlaySupport(); err != nil {
		fmt.Printf("❌ Overlay support check failed: %v\n", err)
		return
	}
	
	// Get filesystem info
	info, err := fm.GetFilesystemInfo(containerFS)
	if err != nil {
		fmt.Printf("❌ Failed to get filesystem info: %v\n", err)
		return
	}
	
	fmt.Printf("Container ID: %s\n", info.ContainerID)
	fmt.Printf("Image: %s\n", info.Image)
	fmt.Printf("Merged Path: %s\n", info.MergedPath)
	fmt.Printf("Layers: %d\n", info.Layers)
	fmt.Printf("Total Size: %s\n", formatBytes(info.TotalSize))
	fmt.Printf("Merged Size: %s\n", formatBytes(info.Size))
	
	fmt.Println("\nLayers:")
	for i, layer := range containerFS.Layers {
		fmt.Printf("  %d. %s (%s)\n", i+1, layer.ID, layer.Digest)
	}
}

// FilesystemInfo represents filesystem information
type FilesystemInfo struct {
	ContainerID string `json:"container_id"`
	Image       string `json:"image"`
	MergedPath  string `json:"merged_path"`
	Layers      int    `json:"layers"`
	TotalSize   int64  `json:"total_size"`
	Size        int64  `json:"size"`
}

// formatBytes formats bytes into human readable format
func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
