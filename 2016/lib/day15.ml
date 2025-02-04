let read_entries ic =
  let rec read acc =
    if Scanf.Scanning.end_of_input ic then
      List.rev acc
    else
      Scanf.bscanf ic "Disc #%d has %d positions; at time=0, it is at position %d.\n" (fun _ size offset ->
          read ((size, offset) :: acc))
  in
  read []

let read filename = Scanf.Scanning.open_in filename |> read_entries

(*
 * Quoting from Wikipedia (https://en.wikipedia.org/wiki/Chinese_remainder_theorem):
 *
 * [T]he Chinese remainder theorem states that if one knows the remainders of
 * the Euclidean division of an integer n by several integers, then one can
 * determine uniquely the remainder of the division of n by the product of
 * these integers, under the condition that the divisors are pairwise coprime
 * (no two divisors share a common factor other than 1).
 *
 * https://blog.ocaml.xyz/algorithm/2022/07/01/chinese-remainder-theorem.html
 * has an implementation of this, and that's what I'm copying here. TBH, my
 * understanding of this area of maths is a bit rubbish.
 *)

exception Modular_inverse

let inverse_mod a = function
  | 1 -> 1
  | b ->
      let rec inner a b x0 x1 =
        if a <= 1 then
          x1
        else if b = 0 then
          raise Modular_inverse
        else
          inner b (a mod b) (x1 - (a / b * x0)) x0
      in
      let x = inner a b 0 1 in
      if x < 0 then x + b else x

let chinese_remainder_exn congruences =
  let mtot = List.map snd congruences |> List.fold_left ( * ) 1 in
  let sum =
    List.fold_left
      (fun acc (r, n) ->
        let rest = mtot / n in
        acc + (r * inverse_mod rest n * rest))
      0 congruences
  in
  sum mod mtot

let to_congruence = List.mapi (fun i (size, offset) -> (size - offset - i - 1, size))
let part_one input = to_congruence input |> chinese_remainder_exn |> Printf.printf "Part 1: %d\n%!"
let part_two input = to_congruence (input @ [(11, 0)]) |> chinese_remainder_exn |> Printf.printf "Part 2: %d\n%!"
