# Name
**license** - automatically generate custom licenses for your repositories

# Synopsis
**license** [flags] [license-type]

# Description
**license** generates a copy of a license file and prints it to the standard output where it can be directed to a file. It also accepts an *-o/-output* argument for writing directly to a file.

# Options
**-help**  
: Print help information  
**-list**  
: Print list of available license types  
**-n**, **-name**  
: Full name to use on license (default username if git/h-config or *LICENSE_FULL_NAME* environment-variable  are not set)  
**-o**, **-output**  
: Path to output file (prints to stdout if unspecified)  
**-v**, **-version**  
: Print version  
**-y**, **-year**  
: Year to use on license (defaults to current year)  

# Examples
**license** mit  
**license** -name "Alice L" -year 2013 bsd-3-clause  
**license** -o LICENSE.txt mpl-2.0  

# Exit Values
**0** No Error  
**1** User Input Error  
**2** Internal Error  
