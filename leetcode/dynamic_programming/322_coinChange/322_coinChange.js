/**
 * Created by betty on 9/5/20.
 */
"use strict";
//@ts-ignore
let coinChange = (coins, amount) => {
  let mat = new Array(amount + 1)
  mat[0] = 0
  for (let i = 1; i < amount + 1; i++) {
    mat[i] = amount + 1
  }
  for (let a = 1; a <= amount; a++) {
    for (let j = 0; j < coins.length; j++) {
      let alt = amount + 1;
      if (a >= coins[j]) {
        alt = mat[a - coins[j]] + 1
      }
      mat[a] = Math.min(mat[a], alt);
    }
  }
  return mat[amount] < amount + 1 ? mat[amount] : -1
};