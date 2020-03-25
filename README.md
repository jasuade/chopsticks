# Implementation of the Chopsticks game in GO
Source: https://en.wikipedia.org/wiki/Chopsticks_(hand_game)

Chopsticks is a hand game for (usually) two players in which players extend a number of fingers from each hand and transfer those scores by taking turns to tap one hand against another.

## Rules (Cutoff style)
1. Each player begins with 1 finger on each hand. Any player goes first, and turns proceed clockwise.
2. On your turn, you must either attack or split, but not both.
3. To *attack*, use one of your live hands to strike an opponent's live hand. When you do this, the number of fingers on the opponent's hand will increase by the number of fingers on the hand you used to strike.
4. To *split*, put your own two hands together, and transfer fingers from one hand to the other as desired. You must make a meaningful move (e.g. going [2, 3]–>[3, 2] is prohibited) and you must not kill one of your own hands (e.g. going [1, 1]–>[0, 2] is prohibited).
5. If any hand of any player reaches five fingers (or more), then it is killed and becomes dead. A player may revive their own dead hand using a split, as long as they follow the rules in #4. However, players may not revive opponents' hands using an attack. Therefore, a player with two dead hands is knocked out of the game and **loses**.
6. You win once all your opponents are knocked out of the game (by each having two dead hands).

**Deviation**

5.1. **Roll-over case** If any hand of any player reaches more than five fingers, then subtract five fingers from that hand to keep it alive. In other words, if a 4-finger hand strikes a 2-finger hand, then the 2-finger hand will have 6 fingers. But since it has more than 5 fingers, then subtract 5 fingers. Now the hand that was struck will have only 1 finger, and therefore it is still alive. This is called roll-over.


**NOTES**: 
A chopsticks position can be represented by a *four-digit code [ABCD]*. A and B are the hands (in ascending order of fingers) of the player who is about to take their turn. C and D are the hands of the oponent. Note the ascending order, so that a single distinct position isn't accidentally represented by two codes. e.g. the code [1032] is wrong, should be notated [0123].

There are 625 positions (including redundancies),225 functionally distinct positions, and **204 reachable positions.**, except if you are playing a variant.

Under normal rules, there are a maximum of 14 possible moves:

- Four attacks (A-C, A-D, B-C, B-D)
- Four divisions (02->11, 03->12, 04->13, 04->22)
- Six transfers (13->22, 22->13, 14->23, 23->14, 24->33, 33->24)

However, only 5 or less of these are available on a given turn.

### Configurations:

A file with different fields of configuration should be reding at the beginning.

The first implementation of this game will be done without configuration using the cutoff example. Them implementing the with roll-over, and later on follow the next configuration.

These configurations include;
- Difficulty: easy (the computer execute random allowed movements without aim to win), medium (there is aim to win but with restrictios), hard (always aim to win or tie)
- Palyers: 1 (to play against the machine), 2 (to play between two players), 3 or more (to be decided)
- Other configurations (Modes):
    - Suicide: You are allowed to kill one of your own hands with a split. You can [12-01] --> [03-01].
    - Swaps: If you have two unequal live hands, you may swap them (essentially forfeiting your turn).
    - Sudden Death: You lose when you only have one finger left total. Alternately, each player could begin with three lives, and every time they get down to [01], they lose a life.
    - Meta: If your hands add up to over five, you can combine them, subtract five from the total, and then split up the remainder. For example, [44] adds up to 8. So under Meta rules, you can combine them into 8, which becomes 3, which you could then split into [12]. Therefore you could go from [44] to [12] in a single move. Meta unlocks 2 new possible moves (34-11, 44-12). If playing both Meta and Suicide, then 4 additional moves are unlocked (24-01, 33-01, 34-02, 44-03), for a total maximum of 20 possible moves.
    - Logan Clause: You are allowed to suicide and swap, but only if you do both at the same time (i.e. swap your dead hand for your live one).
    - **Cutoff: If a hand gets above 5 fingers, it is dead.**
    - Zombies: In 3 or more players, if a player is knocked out, then he is permanently reduced to 1 finger on 1 hand. On his turn, he may attack, but he may not split or be attacked.
    - Halvesies: Splitting is only allowed if you are dividing an even number into two equal halves.
    - Stumps: If you are at [01], you are allowed to split to [0.5 0.5].
    - Suns: Both players start with a 4 in each of their hands ([4444]). *It has to be played andatory with Roll-over* . This is a position that is unreachable in normal gameplay.
