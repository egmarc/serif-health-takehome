# Take-home Write-Up

Important note: **This solution is for the MARCH 2024 file.** I initially began this assignment last night, Feb 29th (leap year!)
but it seems that Anthem was updating the file so the February file was unavailable (see *assets/access_denied.jpg* 
as proof). Their [website](https://www.anthem.com/machine-readable-file/search/) mentioned that they were down for maintenance 
and to try again tomorrow (March 1st) - see *assets/refresh.jpg* as proof.

### Bootstrapping

Clone this repo (git@github.com:egmarc/serif-health-takehome.git)

You will need to have [Go installed](https://go.dev/doc/install).

If you are using a Mac you will also need to [install Xcode](https://apps.apple.com/us/app/xcode/id497799835?mt=12).

### Solution Walkthrough

The list of URLs is under the filename **filtered_locations.txt.gz** found in the root dir of this project.

Overall this code challenge took me roughly ~3 hours. Writing the actual code was the easiest/fastest portion. What took 
the most time was trying to figure out exactly what the take-home was asking for. I attribute that to unfamiliarity with 
insurance/medical data. The schema and sample files were somewhat helpful, but it wasn't until I downloaded the file and 
opened it (all 27GB+) via Hex Fiend where I got a much better understanding of the data and what you were looking for. 
Running the script took a long time as well and contributed to the 3 hours. I didn't time it, but it was in the vicinity 
of 30+ min.

Matthew made it clear he wasn't expecting time spent on fancy bells and whistles, so I kept things simple - stream the 
file, unzip it, and then use the json decoder to retrieve and iterate over the **in_network_files** JSON array. The 
description was vital in retrieving New York PPO links. Everything was under Highmark. Looking online it seems Anthem 
and Highmark are two separate companies under the BCBS federation? If the string inside the *"description"* field 
included *"New York"* and *"PPO"* then I retrieved the URL from the *"location"* field and saved it to filtered_locations.txt.