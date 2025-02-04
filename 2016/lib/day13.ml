let ord_to_magic (x, y) key = (x * x) + (3 * x) + (2 * x * y) + y + (y * y) + key

(* Kernighan's method *)
let count_set_bits n =
  let total = ref 0
  and remaining = ref n in
  while !remaining <> 0 do
    remaining := !remaining land (!remaining - 1);
    total := !total + 1
  done;
  !total

type cell =
  | Wall
  | Empty

type status =
  | Visited
  | Unvisited

let generate_cell key (x, y) = if (ord_to_magic (x, y) key |> count_set_bits) land 1 = 1 then Wall else Empty

let follow visit generator =
  let map = Hashtbl.create 100 in
  let check (x, y) =
    if x < 0 || y < 0 then
      (Visited, Wall)
    else
      match Hashtbl.find_opt map (x, y) with
      | Some kind -> (Visited, kind)
      | None ->
          let kind = generator (x, y) in
          Hashtbl.add map (x, y) kind;
          (Unvisited, kind)
  in
  let paths = Queue.create () in
  let push (x, y) distance =
    match check (x, y) with
    | (Unvisited, Empty) -> Queue.push ((x, y), distance) paths
    | (_, _) -> ()
  in
  let rec loop () =
    if Queue.is_empty paths then
      0
    else
      let ((x, y), distance) = Queue.take paths in
      if visit (x, y) distance then
        distance
      else (
        push (x + 1, y) (distance + 1);
        push (x - 1, y) (distance + 1);
        push (x, y + 1) (distance + 1);
        push (x, y - 1) (distance + 1);
        loop ())
  in
  Queue.push ((1, 1), 0) paths;
  loop ()

let read filename = Io.read_all filename |> String.trim |> int_of_string
let part_one input = follow (fun ord _ -> ord = (31, 39)) (generate_cell input) |> Printf.printf "Part 1: %d\n%!"

let part_two input =
  let visited = ref 0 in
  follow
    (fun _ distance ->
      visited := !visited + 1;
      distance > 50)
    (generate_cell input)
  |> ignore;
  Printf.printf "Part 2: %d\n%!" !visited
