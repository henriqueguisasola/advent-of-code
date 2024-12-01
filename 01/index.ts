import { data } from "./data";

const part1 = () => {
  const list1 = [];
  const list2 = [];
  let sumDiff = 0;

  const pairs = data.split("\n");
  pairs.forEach((pair) => {
    const [first, second] = pair.split("   ");
    list1.push(parseInt(first));
    list2.push(parseInt(second));
  });

  list1.sort();
  list2.sort();

  list1.forEach((first, index) => {
    sumDiff += Math.abs(first - list2[index]);
  });

  console.log(sumDiff);
};

const part2 = () => {
  const list1 = [];
  const list2 = {};
  let similarity = 0;

  const pairs = data.split("\n");
  pairs.forEach((pair) => {
    const [first, second] = pair.split("   ");
    list1.push(parseInt(first));

    list2[second] ? (list2[second] += 1) : (list2[second] = 1);
  });
  list1.forEach((first) => {
    similarity += first * (list2[first] || 0);
  });

  console.log(similarity);
};

// part1();
part2();
