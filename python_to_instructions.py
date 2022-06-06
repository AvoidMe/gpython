"""
Due to constantly changing format of pyc files, I decide to write this script,
   which will compile single python file and dump it in a json-format
"""
import argparse
import dis
import json


def parse_args():
    parser = argparse.ArgumentParser(
        description="Tool for extracting python bytecode from sources"
    )
    parser.add_argument("source_file", help="Python file to compile")
    parser.add_argument("output_file", help=".json file with dumped bytecode")
    return parser.parse_args()


def main():
    args = parse_args()
    source_content = open(args.source_file, "r").read()
    bytecode = []
    for instruction in dis.get_instructions(source_content):
        # print(instruction)
        bytecode.append(
            {
                "opcode": instruction.opcode,
                "opname": instruction.opname,
                "arg": instruction.arg,
                "argval": instruction.argval,
                "offset": instruction.offset,
                "is_jump_target": instruction.is_jump_target,
            }
        )
    json.dump(bytecode, open(args.output_file, "w"))


if __name__ == "__main__":
    main()
