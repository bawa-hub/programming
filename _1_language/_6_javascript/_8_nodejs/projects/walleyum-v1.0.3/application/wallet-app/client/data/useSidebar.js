import create from 'zustand';

const useSidebar = create((set) => ({
  visible: false,
  toggle: () => set((state) => ({ visible: !state.visible })),
}));

export default useSidebar;
