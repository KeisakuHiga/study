function getTriangle(args) {
  if (args.base === undefined) { args.base = 1 }
  if (args.hight === undefined) { args.hight === 1 }
  return args.base * args.hight / 2
}

console.log({ base: 5, hight: 10})
