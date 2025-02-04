let rec gcd a b =
  if b <> 0 then
    gcd b (a mod b)
  else
    a

let lcm xs =
  let rec loop acc xs =
    match xs with
    | x :: xs' -> loop (acc * x / gcd acc x) xs'
    | [] -> acc
  in
  loop 1 xs
