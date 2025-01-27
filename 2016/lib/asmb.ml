(*
 * Assembunny interpreter
 *)

type reg =
  | A
  | B
  | C
  | D

type bytecode =
  | CpyI of int * reg
  | CpyR of reg * reg
  | Inc of reg
  | Dec of reg
  | JnzI of int * int
  | JnzR of reg * int

type machine = {
  a : int;
  b : int;
  c : int;
  d : int;
  pc : int;
}

let initialise ?(a = 0) ?(b = 0) ?(c = 0) ?(d = 0) () = { a; b; c; d; pc = 0 }

let execute program initial =
  let set machine value = function
    | A -> { machine with a = value; pc = machine.pc + 1 }
    | B -> { machine with b = value; pc = machine.pc + 1 }
    | C -> { machine with c = value; pc = machine.pc + 1 }
    | D -> { machine with d = value; pc = machine.pc + 1 }
  and get machine = function
    | A -> machine.a
    | B -> machine.b
    | C -> machine.c
    | D -> machine.d
  in
  let rec loop machine =
    if machine.pc >= Array.length program then
      machine
    else
      loop
        (match program.(machine.pc) with
        | CpyI (value, dest) -> set machine value dest
        | CpyR (src, dest) -> set machine (get machine src) dest
        | Inc reg -> set machine (get machine reg + 1) reg
        | Dec reg -> set machine (get machine reg - 1) reg
        | JnzI (value, offset) -> { machine with pc = (machine.pc + if value <> 0 then offset else 1) }
        | JnzR (reg, offset) -> { machine with pc = (machine.pc + if get machine reg <> 0 then offset else 1) })
  in
  loop initial
