# threaterator
Threat Idea Generator that returns potential threat based on known vulnerabilities for selected software

## Business Objective
Provide a list of potential threats based on software or programming language entered to assist with threat modeling. 

## Requirements
General objective for the threaterator is to iterate known vulnerability information for plausible threats
<ul>
  <li>Run from terminal or command prompt</li>
  <li>Search for key terms</li>
  <li>Use locally stored CVE CSV from MITRE to find potential threats</li>
  <li>Search CSV file for key terms</li>
  <li>Return results of generated potential threats to consider while threat modeling system</li>
  <ul><li>First Display results by Threat type</li>
  <li>Then show potential threats found </li>
  <li>Then show userTerm + key terms</li> //is this needed?
  <li>followed by full description for context</li></ul>
</ul>

## Use case
So far there really is just this one use case
<ol>
<li>User runs threaterator</li>
<li>User enters keyword(userTerm) for software or programming language</li>
<li>Print "Finding threats related to (userTerm)</li>
<li>Threaterator searches Description column of MITRE CVE</li>
<li>If (userTerm) is found, it searches for key terms from Parameters spreadsheet</li>
<li>True = match found, print results</li>
<li>False = match not found, print all results with (userTerm)</li>
</ol>

##Misuse Case #1
User misuses threaterator by fuzzing<br>
This will likely break it, but it will be running locally so risk accepted, lol<br>
Mitigation - Sanitize user data and/or provide user with an prepared list of key software terms

## Spreadsheet Info
MITRE CVE - search this one
<ul>Rows:
<li>1 = Version</li>
<li>2 = Date</li>
<li>3 = Headings (Name, Status, Description, References, Phase, Votes, Comments)</li>
<li>4-10 = Unimportant for threaterator</li>
<li>11-fin = CVE information</li>
</ul>

<ul>Columns:
<li>A = CVE# (aka Name), potential future use when linked to Description Info</li>
<li>B = Status, inconsequential for threaterator use</li>
<li>C = Description, search this column for information</li>
<li>D = Reference, may be useful for potential future use</li>
<li>E = Phase, inconsequential for threaterator use</li>
<li>F = Votes, inconsequential for threaterator use</li>
<li>G = Notes, inconsequential for threaterator use</li>
</ul>

Potential Threats.xls
<ul>Rows
<li>1 = Headers (Threat Type, Threat, Key Terms, Key Words), don't include in search</li>
<li>2-fin = Use for key parameters (Key Terms & Key Words columns) as well as info to return (Threat Type, Threat)</li>
</ul>
<ul>Columns
<li>A = Threat Type, STRIDE+++ listed as options</li>
<li>B = Threat, potential threat</li>
<li>C = Key Terms, used to search MITRE CVE</li>
<li>D = Key Words, used with Key Terms to search MITRE CVE</li>
</ul>


## Future Features
<ol>
  <li>Retrieve CSV file from https://www.cve.org/Downloads and store this file locally</li>
  <li>Include dropdown list for software options</li>
  <li>Retrieve info from other security databases, APIs, or websites</li>
  <li>Retrieve info from software specific sources</li>
  <li>Return results in CSV format</li>
  <li>Return results with more information concerning that specific threat</li>
</ol>

