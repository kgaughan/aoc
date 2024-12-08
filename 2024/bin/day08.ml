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

let mark_antinodes width height towers =
  let antinodes = Array.make_matrix height width false in
  let set_antinode x y =
    if x >= 0 && y >= 0 && x < width && y < height then
      antinodes.(y).(x) <- true
  in
  let loop locations =
    let last = Array.length locations - 1 in
    for i = 0 to last do
      let (x1, y1) = locations.(i) in
      for j = i + 1 to last do
        let (x2, y2) = locations.(j) in
        let (dx, dy) = (x2 - x1, y2 - y1) in
        set_antinode (x1 - dx) (y1 - dy);
        set_antinode (x2 + dx) (y2 + dy)
      done
    done
  in
  Hashtbl.iter (fun _ locations -> loop (Array.of_list locations)) towers;
  antinodes

let count_antinodes antinodes =
  Array.fold_left (fun acc row -> Array.fold_left (fun acc b -> if b then acc + 1 else acc) acc row) 0 antinodes

let _ =
  let grid = read_input "input/day08-sample.txt" in
  let height = Array.length grid
  and width = String.length grid.(0)
  and towers = find_towers grid in
  let part1 = mark_antinodes width height towers |> count_antinodes in
  Printf.printf "Part 1: %d\n" part1
