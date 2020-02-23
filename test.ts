// create subprocess
const p = Deno.run({
    args: ["$USER/projects/sophos-cli/target/debug/sophos-cli", "9919103015", "052019UH", "login"]
  });
  
  // await its completion
  await p.status();