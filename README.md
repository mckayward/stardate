# Challenge 063: Stardate [![CircleCI](https://circleci.com/gh/mckayward/stardate.svg?style=shield)](https://circleci.com/gh/mckayward/stardate)

In preparation for interplanetary space travel and humans living on planets other than earth, we are going to need a new system for expressing dates. Star Trek kinda already did this in the future for us, with the stardate system. We’re going to need some functions to translate earth dates into stardates (and vice versa). Your challenge is to write a function that converts and earth date (ie. 2016-10-05) into a stardate.

[This web page](http://trekguide.com/Stardates.htm) provides details on how stardates are calculated. Here is a summary of what you need to know:

Stardate 00000.0 started on Friday, July 5, 2318, around noon (Starfleet Command time).
There are 918.23186 Stardates per earth year.
(0.397766856 day to 1.0 Stardate, or 1.0 Stardate to 34,367.0564 seconds)

For dates before July 5, 2318, use negative numbers to represent Stardates (None of this YYMM.DD nonsense). If you want to check your logic, target The Next Generation Stardate calculator here: http://trekguide.com/Stardates.htm#TNGcalculator.
