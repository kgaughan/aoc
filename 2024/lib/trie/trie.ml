type t = {
  mutable is_leaf : bool;
  mutable branches : (char * t) list;
}

let create () = { branches = []; is_leaf = false }

let add s t =
  let l = String.length s in
  let rec insert i t =
    let is_leaf = i = l - 1 in
    match List.assoc_opt s.[i] t.branches with
    | Some t' ->
        if is_leaf then
          t'.is_leaf <- true
        else
          insert (i + 1) t'
    | None ->
        let new_entry = { is_leaf; branches = [] } in
        if not is_leaf then insert (i + 1) new_entry;
        t.branches <- (s.[i], new_entry) :: t.branches
  in
  insert 0 t

let find_prefixes s t =
  let rec loop i acc t =
    let acc' = if t.is_leaf then String.sub s 0 i :: acc else acc in
    if i = String.length s then
      acc'
    else
      List.fold_left (fun acc (ch, t) -> if ch = s.[i] then loop (i + 1) acc t else acc) acc' t.branches
  in
  let result = loop 0 [] t in
  result
