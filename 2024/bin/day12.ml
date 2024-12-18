type plot =
  | Unscanned of char
  | Scanned of char

let directions = [(0, -1); (0, 1); (1, 0); (-1, 0)]

let read_input path =
  Utils.read_lines path (Utils.array_of_string (fun ch -> Unscanned ch) (Unscanned ' ')) |> Array.of_list

(* The plan for this one is to first trace the edge of each region to determine
   the perimeter. Once that's obtained, the rest of the plots in the region
   can be found to determine its area. *)

let trace_region x y ch map =
  let area = ref 0
  and perimeter = ref 0
  and corners = ref 0
  and pending = Queue.create () in
  let in_bounds x y = x >= 0 && y >= 0 && x < Array.length map.(0) && y < Array.length map in
  let is_perimeter x y =
    if in_bounds x y then
      match map.(y).(x) with
      | Unscanned ch' -> ch' <> ch
      | Scanned ch' -> ch' <> ch
    else
      true
  in
  let try_move x y =
    if in_bounds x y then
      match map.(y).(x) with
      | Unscanned ch' when ch' = ch -> Some (x, y)
      | _ -> None
    else
      None
  in
  let count_corners x y =
    let np = is_perimeter x (y - 1)
    and sp = is_perimeter x (y + 1)
    and wp = is_perimeter (x - 1) y
    and ep = is_perimeter (x + 1) y
    and nwc = is_perimeter (x - 1) (y - 1)
    and nec = is_perimeter (x + 1) (y - 1)
    and swc = is_perimeter (x - 1) (y + 1)
    and sec = is_perimeter (x + 1) (y + 1) in
    List.fold_left
      (fun acc b -> acc + if b then 1 else 0)
      0
      [
        (* convex corners *)
        np && wp;
        np && ep;
        sp && wp;
        sp && ep;
        (* concave corners *)
        nwc && not (np || wp);
        nec && not (np || ep);
        swc && not (sp || wp);
        sec && not (sp || ep);
      ]
  in
  let push (x, y) =
    Queue.add (x, y) pending;
    map.(y).(x) <- Scanned ch;
    area := !area + 1;
    perimeter :=
      List.fold_left (fun acc (dx, dy) -> acc + if is_perimeter (x + dx) (y + dy) then 1 else 0) !perimeter directions;
    corners := !corners + count_corners x y
  in
  push (x, y);
  let rec loop () =
    if Queue.length pending = 0 then
      (!area, !perimeter, !corners)
    else
      let (x', y') = Queue.take pending in
      List.iter (fun (dx, dy) -> try_move (x' + dx) (y' + dy) |> Option.iter push) directions;
      loop ()
  in
  loop ()

let discover_regions grid =
  let height = Array.length grid
  and width = Array.length grid.(0)
  and regions = Hashtbl.create 256
  and id = ref 0 in
  for y = 0 to height - 1 do
    for x = 0 to width - 1 do
      match grid.(y).(x) with
      | Unscanned ch ->
          Hashtbl.add regions !id (trace_region x y ch grid);
          id := !id + 1
      | _ -> ()
    done
  done;
  regions

let _ =
  let grid = read_input "input/day12.txt" in
  let regions = discover_regions grid in
  let part1 = Hashtbl.fold (fun _ (area, perimeter, _) acc -> acc + (area * perimeter)) regions 0 in
  let part2 = Hashtbl.fold (fun _ (area, _, corners) acc -> acc + (area * corners)) regions 0 in
  Printf.printf "Part 1: %d; Part 2: %d\n" part1 part2
