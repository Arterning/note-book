import os
import shutil


def process_files(directory):
    for filename in os.listdir(directory):
        if "-" in filename:
            parts = filename.split("-")
            if len(parts) > 1:
                prefix = parts[0]
                new_directory = os.path.join(directory, prefix)
                if not os.path.exists(new_directory):
                    os.makedirs(new_directory)

                source_path = os.path.join(directory, filename)
                destination_path = os.path.join(new_directory, filename)

                shutil.move(source_path, destination_path)
                print(f"Moved: {source_path} -> {destination_path}")


if __name__ == "__main__":
    current_directory = os.getcwd()
    process_files(current_directory)