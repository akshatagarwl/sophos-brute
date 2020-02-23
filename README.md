# Sophos Password Vulnerabilty Checker

Most of us don't bother to get the password (that includes the people who made this) given to us by the college. Use this tool to find out if your password is vulnerable or not.

P.S. Ofc your password would still be vulnerable even if you get it changed. Changing it would just mean that there is lesser chance of someone like us using it

This is was made in [Deno]. You need to install it in order to use it.

Before you can start using it open the shhh.ts file and edit
    - line 1 : Add your enrollment number
    - line 2 : Add your password. This has to be added because Sophos doesn't let you connect for a while after 4 consecutive unsuccesful attempts
    - line 4 : Add the enrollment number of the person whose password you want to find

You still have to manually check and see the output where it was successful. 

To use type `deno --allow-run shhh.ts`

TODO:
- [ ] Prompt the user for username and password
- [ ] Not print every trial
- [ ] Detect the correct trial, ouput the password and stop the program 

<!-- links -->
[Deno]: https://deno.land/