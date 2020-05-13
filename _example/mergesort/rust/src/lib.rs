#[inline]
pub fn fibonacci(n: u64) -> u64 {
    match n {
        0 => 1,
        1 => 1,
        n => fibonacci(n - 1) + fibonacci(n - 2),
    }
}

pub fn mergesort(src: &[i64]) -> Vec<i64> {
    if src.len() == 1 {
        return vec![src[0]];
    }
    if src.len() == 2 {
        return if src[0] > src[1] {
            vec![src[1], src[0]]
        } else {
            vec![src[0], src[1]]
        };
    }

    let mid = src.len() / 2;
    let l = mergesort(&src[..mid]);
    let r = mergesort(&src[mid..]);
    if l[mid - 1] <= r[0] {
        // skip merge
        let mut v = Vec::with_capacity(src.len());
        v.extend_from_slice(&l);
        v.extend_from_slice(&r);
        return v;
    }
    return merge(&l, &r);
}

pub fn merge(a: &[i64], b: &[i64]) -> Vec<i64> {
    let la = a.len();
    let lb = b.len();
    let lc = la + lb;
    let mut v = Vec::with_capacity(lc);

    let mut i = 0;
    let mut ia = 0;
    let mut ib = 0;
    while i < lc {
        if ia == la || ib == lb {
            break;
        }
        if a[ia] <= b[ib] {
            v.push(a[ia]);
            ia += 1;
        } else {
            v.push(b[ib]);
            ib += 1;
        }
        i += 1;
    }

    // copy the rest
    if i == lc {
        return v;
    }
    if ia < la {
        v.extend_from_slice(&a[ia..]);
    }
    if ib < lb {
        v.extend_from_slice(&b[ib..]);
    }
    v
}

#[cfg(test)]
mod tests {
    use rand::{thread_rng, Rng};
    use crate::mergesort;

    #[test]
    fn it_is_correct() {
        let mut rng = thread_rng();
        let mut unsorted = Vec::new();
        let size = 1000;
        let mut i = 0;
        while i < size {
            let x: i64 = rng.gen();
            unsorted.push(x);
            i += 1;
        }
        let sorted = mergesort(&unsorted);
        assert_eq!(sorted.len(), unsorted.len());
        // NOTE: the builtin sort is merge sort like timsort
        unsorted.sort();
        assert_eq!(sorted, unsorted);
    }
}
