# Sophos Password Vulnerabilty Checker
Disclaimer: This is just for education purpose and I do not intend to steal any information.

This is was made in [Deno]. You need to install it in order to use it. This uses a modified version of [sophos-cli](https://github.com/ryzokuken/sophos-cli) (with an added logout option). The binary is shipped with the repository. Currently works for most first year students i.e having enrollment number `19*` or `9919*`

Before you can start using it open the script.ts file and edit
- line 1 : Add your enrollment number
- line 2 : Add your password. This has to be added because Sophos doesn't let you connect for a while after 4 consecutive unsuccesful attempts
- line 4 : Add the enrollment number of the person whose password you want to find

You still have to manually check and see the output where it was successful. 

To use type `deno --allow-run script.ts`

TODO:
- [ ] Prompt the user for username and password
- [ ] Not print every trial
- [ ] Detect the correct trial, ouput the password and stop the program
- [ ] Put the passwords in a different file

## Collaborators
- [DelusionalOptimist](https://github.com/DelusionalOptimist)
- [humancalico](https://github.com/humancalico)
<!-- links -->
[Deno]: https://deno.land/
