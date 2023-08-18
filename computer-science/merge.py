import os
import sys


def merge_markdown_files(directory, output_filename):
    markdown_files = [file for file in os.listdir(directory) if file.endswith(".md")]

    output_path = os.path.join(directory, output_filename)

    with open(output_path, "w", encoding="utf-8") as output_file:
        for markdown_file in markdown_files:
            file_path = os.path.join(directory, markdown_file)
            with open(file_path, "r", encoding="utf-8") as input_file:
                output_file.write("# " + markdown_file + "\n\n")
                output_file.write(input_file.read() + "\n\n")

    print(f"Markdown files from {directory} merged into {output_path}")


if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: python script.py <directory> <output_filename>")
        input_directory = "tool/git"
        output_filename = "merged.md"
        merge_markdown_files(input_directory, output_filename)
    else:
        # input_directory = sys.argv[1]
        # output_filename = sys.argv[2]
        input_directory = "思考"
        output_filename = "merged.md"
        merge_markdown_files(input_directory, output_filename)
