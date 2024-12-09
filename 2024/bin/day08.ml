let read_input path = In_channel.with_open_text path Utils.input_lines |> Array.of_list

let find_towers grid =
  let towers = Hashtbl.create 62 in
  let search_vertically y row =
    let add x c =
      if c <> '.' then
        match Hashtbl.find_opt towers c with
        | Some tl -> Hashtbl.replace towers c ((x, y) :: tl)
        | None -> Hashtbl.add towers c [(x, y)]
    in
    String.iteri add row
  in
  Array.iteri search_vertically grid;
  towers

let check_pairs marker locations =
  let last = Array.length locations - 1 in
  for i = 0 to last do
    for j = i + 1 to last do
      marker locations.(i) locations.(j)
    done
  done

let mark_antinodes width height towers =
  let antinodes = Array.make_matrix height width false in
  let set_antinode x y =
    if x >= 0 && y >= 0 && x < width && y < height then
      antinodes.(y).(x) <- true
  in
  let mark (x1, y1) (x2, y2) =
    let (dx, dy) = (x2 - x1, y2 - y1) in
    set_antinode (x1 - dx) (y1 - dy);
    set_antinode (x2 + dx) (y2 + dy)
  in
  Hashtbl.iter (fun _ locations -> check_pairs mark (Array.of_list locations)) towers;
  antinodes

let mark_antinodes_repeating width height towers =
  let antinodes = Array.make_matrix height width false in
  let in_bounds x y = x >= 0 && y >= 0 && x < width && y < height in
  (* There are better ways of doing this, but I just can't be bothered *)
  let mark (x1, y1) (x2, y2) =
    let dx = x1 - x2
    and dy = y1 - y2
    and x' = ref x1
    and y' = ref y1 in
    while in_bounds !x' !y' do
      antinodes.(!y').(!x') <- true;
      x' := !x' - dx;
      y' := !y' - dy
    done;
    x' := x1 + dx;
    y' := y1 + dy;
    while in_bounds !x' !y' do
      antinodes.(!y').(!x') <- true;
      x' := !x' + dx;
      y' := !y' + dy
    done
  in
  Hashtbl.iter (fun _ locations -> check_pairs mark (Array.of_list locations)) towers;
  antinodes

let count_antinodes antinodes =
  Array.fold_left (fun acc row -> Array.fold_left (fun acc b -> if b then acc + 1 else acc) acc row) 0 antinodes

let _ =
  let grid = read_input "input/day08.txt" in
  let height = Array.length grid
  and width = String.length grid.(0)
  and towers = find_towers grid in
  let part1 = mark_antinodes width height towers |> count_antinodes in
  let part2 = mark_antinodes_repeating width height towers |> count_antinodes in
  Printf.printf "Part 1: %d; Part 2: %d\n" part1 part2
