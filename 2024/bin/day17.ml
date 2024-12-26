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
      output
    else
      let (machine', output') = simulate_once machine output in
      loop machine' output'
  in
  loop machine []

let find_quine machine =
  let check_trailing lst =
    let rec loop n = function
      | hd :: tl when hd = machine.program.(n) -> loop (n - 1) tl
      | [] -> true
      | _ -> false
    in
    loop (Array.length machine.program - 1) lst
  in
  let rec loop offset acc =
    let rec loop_candidate c =
      let attempt = (acc * 8) + c in
      if simulate { machine with reg_a = attempt } |> check_trailing then
        if offset = 0 then
          Some attempt
        else
          match loop (offset - 1) attempt with
          | Some n -> Some n
          | None -> if c < 7 then loop_candidate (c + 1) else None
      else if c < 7 then
        loop_candidate (c + 1)
      else
        None
    in
    loop_candidate 0
  in
  loop (Array.length machine.program - 1) 0 |> Option.get

let _ =
  let machine = read_file "input/day17.txt" in
  let part1 = Utils.time "part 1" (fun () -> simulate machine |> List.rev |> Utils.int_concat) in
  let part2 = Utils.time "part 2" (fun () -> find_quine machine) in
  Printf.printf "Part 1: %s; Part 2: %d\n" part1 part2
