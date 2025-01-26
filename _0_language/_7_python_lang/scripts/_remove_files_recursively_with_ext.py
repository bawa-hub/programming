import os

def remove_files_with_extension_recursively(folder_path, extension):
    for root, dirs, files in os.walk(folder_path):
        for filename in files:
            if filename.endswith(extension):
                file_path = os.path.join(root, filename)
                os.remove(file_path)
                print(f"Removed: {file_path}")
                

# Get the current working directory
current_directory = os.getcwd()

# # Get the base name of the current directory (i.e., the last part of the path)
# folder_name = os.path.basename(current_directory)
                

# # Example: Remove all '.txt' files in the 'my_folder' directory and its subdirectories
folder_path = current_directory + '/' + '_7_django'
# folder_path = "/home/bawa/Learn/language/python_lang/_7_django/projects/student_management_system/student_management_system/html-template"
extension_to_remove = '.Identifier'
remove_files_with_extension_recursively(folder_path, extension_to_remove)
print("ho gyo")

