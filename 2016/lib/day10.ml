let read path = Day10_parser.main Day10_lexer.tokens (Lexing.from_channel (open_in path))
let append tbl k v = Hashtbl.replace tbl k (v :: (Hashtbl.find_opt tbl k |> Option.value ~default:[]))

let process input =
  let bots = Hashtbl.create 1000
  and rules = Hashtbl.create 1000
  and received = Hashtbl.create 1000
  and outputs = Hashtbl.create 1000 in
  List.iter
    (function
      | Day10_aux.Receive (value, bot) -> append bots bot value
      | Day10_aux.Give (bot, low, high) -> Hashtbl.add rules bot (low, high))
    input;
  let rec distribute bot =
    let receive value = function
      | Day10_aux.Bot bot ->
          append received bot value;
          append bots bot value;
          distribute bot
      | Day10_aux.Output output -> Hashtbl.add outputs output value
    in
    match Hashtbl.find bots bot with
    | [a; b] ->
        let (low, high) = Hashtbl.find rules bot in
        Hashtbl.replace bots bot [];
        receive (min a b) low;
        receive (max a b) high
    | _ -> ()
  in
  Hashtbl.to_seq_keys bots |> Seq.iter distribute;
  (received, outputs)

let part_one input =
  let (received, _) = process input in
  let check bot =
    let values = Hashtbl.find received bot in
    List.mem 17 values && List.mem 61 values
  in
  received |> Hashtbl.to_seq_keys |> Seq.filter check |> List.of_seq |> List.hd |> Printf.printf "Part 1: %d\n%!"

let part_two input =
  let (_, outputs) = process input in
  let value = Hashtbl.find outputs in
  Printf.printf "Part 2: %d\n%!" (value 0 * value 1 * value 2)
