struct Solution;

const MOD: i64 = 1_000_000_007;

impl Solution {
    fn factorize(mut k: i64) -> Vec<i64> {
        let mut exponents = vec![];
        let mut d = 2;
        while d * d <= k {
            if k % d == 0 {
                let mut exp = 0;
                while k % d == 0 {
                    exp += 1;
                    k /= d;
                }
                exponents.push(exp);
            }
            d += 1;
        }
        if k > 1 {
            exponents.push(1); // remaining prime factor with exponent 1
        }
        exponents
    }

    fn mod_pow(mut base: i64, mut exp: i64, modulus: i64) -> i64 {
        let mut result = 1;
        base %= modulus;
        while exp > 0 {
            if exp % 2 == 1 {
                result = result * base % modulus;
            }
            exp /= 2;
            base = base * base % modulus;
        }
        result
    }

    fn mod_inv(a: i64, modulus: i64) -> i64 {
        Self::mod_pow(a, modulus - 2, modulus)
    }

    fn comb(a: i64, n: i64) -> i64 {
        if a == 0 {
            return 1;
        }
        let mut numerator = 1i64;
        let mut denominator = 1i64;
        for i in 0..a {
            numerator = numerator * ((n + i) % MOD) % MOD;
            denominator = denominator * (i + 1) % MOD;
        }
        numerator * Self::mod_inv(denominator, MOD) % MOD
    }

    pub fn ways_to_fill_array(queries: Vec<Vec<i32>>) -> Vec<i32> {
        queries
            .iter()
            .map(|q| {
                let n = q[0] as i64;
                let k = q[1] as i64;
                let exponents = Self::factorize(k);
                let ans = exponents
                    .iter()
                    .fold(1i64, |acc, &a| acc * Self::comb(a, n) % MOD);
                ans as i32
            })
            .collect()
    }
}
