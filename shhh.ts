const username = "9919103015";
let password = "052019UH";

function login() {
  // create subprocess
  const p = Deno.run({
    args: [
      "/home/akshat/projects/sophos-cli/target/debug/sophos-cli",
      username,
      password,
      "login"
    ]
  });
  // await its completion
  await p.status();
}

function logout() {
  // create subprocess
  const p = Deno.run({
    args: [
      "/home/akshat/projects/sophos-cli/target/debug/sophos-cli",
      username,
      password,
      "logout"
    ]
  });
  // await its completion
  await p.status();
}

let list = [
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  "",
  ""
];
