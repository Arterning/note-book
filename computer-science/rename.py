import os


# os.walk 返回三元组，包含当前目录路径、子目录列表和文件列表。这使得你可以轻松地遍历整个目录树

def change_extension(directory, old_ext, new_ext):
    for root, _, files in os.walk(directory):
        for filename in files:
            if filename.endswith(old_ext):
                old_path = os.path.join(root, filename)
                new_filename = filename[:-len(old_ext)] + new_ext
                new_path = os.path.join(root, new_filename)

                os.rename(old_path, new_path)
                print(f'Renamed: {old_path} -> {new_path}')


if __name__ == "__main__":
    current_directory = os.getcwd()
    old_extension = ".js"  # 要被替换的旧后缀
    new_extension = ".md"  # 新后缀

    change_extension(current_directory, old_extension, new_extension)
