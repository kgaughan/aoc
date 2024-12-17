type machine = {
  reg_a : int;
  reg_b : int;
  reg_c : int;
  ic : int;
  program : int array;
}

let read_file path =
  let lines = Utils.read_lines path Fun.id |> Array.of_list in
  let get_data s = List.nth (String.split_on_char ':' s) 1 |> String.trim in
  {
    ic = 0;
    reg_a = lines.(0) |> get_data |> int_of_string;
    reg_b = lines.(1) |> get_data |> int_of_string;
    reg_c = lines.(2) |> get_data |> int_of_string;
    program = lines.(4) |> get_data |> String.split_on_char ',' |> List.map int_of_string |> Array.of_list;
  }

let simulate_once machine output =
  let combo = function
    | n when n < 4 -> n
    | 4 -> machine.reg_a
    | 5 -> machine.reg_b
    | 6 -> machine.reg_c
    | _ -> raise (Invalid_argument "Bad combo operand encountered")
  in
  let division operand = machine.reg_a / (1 lsl combo operand) in
  let opcode = machine.program.(machine.ic)
  and operand = machine.program.(machine.ic + 1) in
  match opcode with
  | 0 -> ({ machine with reg_a = division operand; ic = machine.ic + 2 }, output)
  | 1 -> ({ machine with reg_b = machine.reg_b lxor operand; ic = machine.ic + 2 }, output)
  | 2 -> ({ machine with reg_b = 7 land combo operand; ic = machine.ic + 2 }, output)
  | 3 -> ({ machine with ic = (if machine.reg_a = 0 then machine.ic + 2 else operand) }, output)
  | 4 -> ({ machine with reg_b = machine.reg_b lxor machine.reg_c; ic = machine.ic + 2 }, output)
  | 5 -> ({ machine with ic = machine.ic + 2 }, (7 land combo operand) :: output)
  | 6 -> ({ machine with reg_b = division operand; ic = machine.ic + 2 }, output)
  | 7 -> ({ machine with reg_c = division operand; ic = machine.ic + 2 }, output)
  | _ -> raise (Invalid_argument (Printf.sprintf "illegal opcode: %d" opcode))

let simulate machine =
  let rec loop machine output =
    if machine.ic = Array.length machine.program then
      List.rev output
    else
      let (machine', output') = simulate_once machine output in
      loop machine' output'
  in
  loop machine []

let find_quine original_machine =
  let rec loop machine output n =
    (* machine exited unexpectedly? *)
    if machine.ic >= Array.length machine.program then
      loop { original_machine with reg_a = n + 1 } [] (n + 1)
    else (* normal flow *)
      let check_for_emit = machine.program.(machine.ic) = 5
      and (machine', output') = simulate_once machine output in
      if check_for_emit then
        let emitted = List.length output' in
        if List.hd output' <> machine.program.(emitted - 1) then
          loop { original_machine with reg_a = n + 1 } [] (n + 1)
        else if emitted = Array.length machine.program then
          n
        else
          loop machine' output' n
      else
        loop machine' output' n
  in
  loop { original_machine with reg_a = 0 } [] 0

let _ =
  let machine = read_file "input/day17.txt" in
  let part1 = simulate machine in
  Printf.printf "Part 1: %s\n" (Utils.int_concat part1);
  let part2 = find_quine machine in
  Printf.printf "Part 2: %d\n" part2
