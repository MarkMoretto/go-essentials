#!/bin/python

from __future__ import annotations

"""
Summary: Command-line script to make pretty Go comments for section headings.
Date: 2022-02-19

Examples
--------
>>> python section_format.py Function Demo 4


// ----------------------- //
// ---- Function Demo ---- //
// ----------------------- //


>>> python section_format.py Function Demo


// --------------------- //
// --- Function Demo --- //
// --------------------- //


"""

from sys import argv

def format_middle(string, buffer = 3):
    pstr = f" {string} "
    fstr = pstr.center(len(pstr) + (2 * buffer), "-")
    return add_sides(fstr)

def format_top_bottom(middle_str: str):
    slen = len(middle_str)
    fstr = repeats("-", slen - 6)
    return add_sides(fstr)

def add_sides(string, side="//"):
    return f"{side} {string} {side}"

def repeats(s: str, n: int) -> str:
    """Return string of string repeated n times."""
    return s * n

def striput(msg: str = None) -> str:
    """Input handler that strips excess whitespace
    and allows for optional message.
    """
    if msg is None:
        return input().strip()
    else:
        return input(msg).strip()

if __name__ == "__main__":
    lines = [""] * 2
    phrase = " "
    side_buffer = None

    # If arguments given...
    if len(argv) > 1:
        args = argv[1:]
        # print("Your args: ", args)
        if len(args) > 1:
            tmp = str(args[-1]).strip()
            if tmp.isdigit():
                side_buffer = int(tmp)
                phrase = " ".join(args[:-1])
            else:
                side_buffer = 3
                phrase = " ".join(args)
        else:
            side_buffer = 3
            phrase = " ".join(args)


    #  ...otherwise, ask for arguments.
    else:
        phrase = striput("Enter phrase: ")

        while side_buffer is None:
            tmp = striput("Enter buffer amount [default: 3]: ")

            if len(tmp) == 0:
                side_buffer = 3
                break
            else:
                try:
                    ntmp = int(tmp)
                    side_buffer = ntmp
                except ValueError:
                    side_buffer = None
 
    lines[1] = format_middle(phrase, side_buffer)
    lines[0] = format_top_bottom(lines[1])
    lines.append(lines[0])

    print("\n")
    print("\n".join(lines))
    print("\n")
