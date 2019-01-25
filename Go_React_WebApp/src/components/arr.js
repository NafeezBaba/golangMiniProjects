const a = [{
    d: 1,
    b: 2
}, {
    d: 2,
    b: 2
}, {
    d: 3,
    b: 2
}, {
    d: 4,
    b: 2
}, {
    d: 5,
    b: 2
}]


const b = a.map((val, ind) => {
    const newd = val.d * 2
    const newb = val.b * 4
    
    return {
        d: newd,
        b: newb
    }
})
console.log("​a", a)
console.log("​b", b)