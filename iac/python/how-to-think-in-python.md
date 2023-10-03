
# Python Installation
## Check whether python3 is installed

```bash
python3 --version
```

## Create and manage projects in Python

To manage projects, you need a virtual environment, packages that will help you develop faster, and a strategy for maintaining packages.  

- Create a Python project.  
- Develop and run your code in isolation on a machine.  
- Use libraries that someone else has written.  
- Restore a project from a list of dependencies.  

What's a virtual environment?

A virtual environment is a self-contained copy of everything needed to run your program. This includes the Python interpreter and any libraries your program needs. By using a virtual environment, you can ensure your program will have access to the correct versions and resources to run correctly.

The basic workflow looks like this:

Create a virtual environment that won't affect the rest of the machine.
Step into the virtual environment, where you specify the Python version and libraries that you need.
Develop your program.

Create a virtual environment ::  
```bash
python -m venv env
```

It will generate following directo structure ::

/env
  /bin
  /include
  /lib
  /src (we can introduce it to put our source code)

Activate virtual environment ::
```bash
source env/bin/activate
```

Deactive virtual environment once done ! ::

To ensure that these packages exist only in your virtual environment, try stepping out of that environment by calling deactivate:
```bash
deactivate
```

What's a package?

One of the main advantages of using external libraries is to speed up the development time of your program. You can fetch such a library on the internet. But by fetching and installing these libraries through a virtual environment, you ensure that you install theses libraries only for the virtual environment and not globally for the entire machine.

Install package/library ::  

List installed pacekages :  
```bash
pip freeze
```

You install a package by using pip. The [list](https://pypi.org/) of available packages.  
```bash
pip install python-dateutil
```


More ways to install a package ::

- Have a set of files on your machine and install from that source:  
```bash
cd <to where the package is on your machine>
python3 -m pip install .
```
- Install from a compressed archive file
```bash
python3 -m pip install package.tar.gz
```

Work with project files ::

How to distribute project to others for collaboration, and how you can manage your dependencies over time. 

Share a project :  

1. Call pip freeze > requirements.txt. This command creates a requirements.txt file with all the packages that the program needs.
2. Create a .gitignore file, and check in your application code and requirements.txt.
3. Check in the code to GitHub.

Consume a project :

1. Fetch the project from GitHub.
2. Create a virtual environment and place yourself in it.
3. Restore the project by using pip install -r requirements.txt. It will look for requirements.txt and fetch and install the packages listed for that file.
4. Run your app.

Manage dependencies ::  

Install a latest version of a package :
```bash
pip install python-dateutil==1.4
```
Findout available version of a package :
```bash
pip install python-dateutil==randomwords  # note 'randomwords' is a tricky !!
pip freeze python-dateutil
```
Upgrade package :
```bash
pip install <name of package> --upgrade
```

Apply an upgrade strategy :
```bash
pip install "python-dateutil==2.7.*" --upgrade # Packages use sermantic versioning Major.Minor.Patch
```

Clean up unused package :
```bash
pip uninstall python-dateutil
pip freeze > requirements.txt  # update the dependency list
pip uninstall -r requirements.txt -y # Remove all packages
```

Use Boolean logic in python :

Booleans are a common type in Python. Their value can only ever be one of two things: true or false. Understanding how to use Boolean values is critical, because you need them to write conditional logic.

Use Strings in python :

Python strings are one of the most important and common types in the language. Learning how to interact with strings, including formatting and replacing text, is an essential skill to develop for working with Python code.

Using quotation marks :
You can enclose Python strings in single, double, or triple quotation marks. Although you can use them interchangeably, it's best to use one type consistently within a project. 

Multiline text
There are a few different ways to define multiple lines of text as a single variable. The most common ways are:

Use a newline character (\n).
Use triple quotation marks (""").
Newline characters separate the text into multiple lines when you print the output:

Some common string methods ::

```bash
heading = "temperatures and facts about the moon"
heading_upper = heading.title()
# Split a String
temperatures_list = temperatures .split()
print("Moon" in "This text will describe facts and challenges with space travel")
# Search & matching for a string
temperatures = """Saturn has a daytime temperature of -170 degrees Celsius, while Mars has -28 Celsius."""
print(temperatures.find("Mars"))
temperatures = """Saturn has a daytime temperature of -170 degrees Celsius, while Mars has -28 Celsius."""
print(temperatures.count("Mars"))
print(temperatures.count("Moon"))
print("The Moon And The Earth".lower())
print("The Moon And The Earth".upper())
# Check Content
temperatures = "Mars Average Temperature: -60 C"
parts = temperatures.split(':')
print(parts)
print(parts[-1])
print("-60".startswith('-'))
# Transform Text
text = "Temperatures on the Moon can vary wildly."
print("temperatures" in text.lower())
moon_facts = ["The Moon is drifting away from the Earth.", "On average, the Moon is moving about 4cm every year."]
print(' '.join(moon_facts))
```
String formating in python :

Besides transforming text and performing basic operations, such as matching and searching, it's essential to format the text when you're presenting information.

- Using the % formatter.  
- Using .format() in a string.  
- Using f-strings.  

Percent sign (%) formatting :  

```bash
mass_percentage = "1/6"
print("On the Moon, you would weigh about %s of your weight on Earth." % mass_percentage)
print("""Both sides of the %s get the same amount of sunlight, but only one side is seen from %s because the %s rotates around its own axis when it orbits %s.""" % ("Moon", "Earth", "Moon", "Earth"))
```

The format() method :  
```bash
mass_percentage = "1/6"
print("On the Moon, you would weigh about {} of your weight on Earth.".format(mass_percentage))
mass_percentage = "1/6"
print("""You are lighter on the {0}, because on the {0} you would weigh about {1} of your weight on Earth.""".format("Moon", mass_percentage))
```

About f-strings :  

As of Python version 3.6, it's possible to use f-strings. These strings look like templates and use the variable names from your code.

```bash
print(f"On the Moon, you would weigh about {mass_percentage} of your weight on Earth.") # f-string with variable
print(f"On the Moon, you would weigh about {round(100/6, 1)}% of your weight on Earth.") # f-string with match expression
subject = "interesting facts about the moon"
heading = f"{subject.title()}"   # f-string with function call
print(heading)
```

Mathematic Operation In Python ::  

The mathematic operators available in Python.
The order of operations.
How to convert strings to numbers.

Python supports four core mathematics operations, addition, substraction, multiplication, division along with others. 

Operators In Python ::

```bash
# Addition
answer = 30 + 12
difference = 30 - 12
product = 30 * 12
quotient = 30 / 12
quotient = 30 / 12  # floor (round-up) division
```
Orders of operations ::

Python honors the order of operation for math.

1. Parentheses
2. Exponents
3. Multiplication and division
4. Addition and subtraction  

```bash
result_1 = 1032 + 26 * 2
print(result_1)
result_2 = 1032 + (26 * 2)
print(result_2)
```

```bash
# Convert String To Number
demo_int = int('215')
print(demo_int)
demo_float = float('215.3')
print(demo_float)
# Absolute Value
print(abs(39 - 16))
print(abs(16 - 39))
# Rounding
print(round(14.5))
# Math Library
round_up = ceil(12.5)
print(round_up)
round_down = floor(12.5)
print(round_down)
```

Python List ::

Lists allow you to store multiple values in a single variable.

- Identify when to use a list.
- Create a list.
- Access a particular item in a list by using indexes.
- Push items to the end of a list.
- Sort and slice a list.

```bash
# Create a list
planets = ["Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"]
# Access list items by index
print("The first planet is", planets[0])
# Determine the length of a list
number_of_planets = len(planets)
print("There are", number_of_planets, "planets in the solar system.")
# Add values to lists
planets.append("Pluto")
# Remove Value from list
planets.pop()  # Goodbye, Pluto
# Use negative index
print("The first planet is", planets[0])
print("The last planet is", planets[-1])
# Find value in a list
jupiter_index = planets.index("Jupiter")
print("Jupiter is the", jupiter_index + 1, "planet from the sun")
```
# Store Numbers in list
```bash
gravity_on_earth = 1.0
gravity_on_the_moon = 0.166
gravity_on_planets = [0.378, 0.907, 1, 0.377, 2.36, 0.916, 0.889, 1.12]
bus_weight = 124054 # in Newtons, on Earth
print("On Earth, a double-decker bus weighs", bus_weight, "N")
print("On Mercury, a double-decker bus weighs", bus_weight * gravity_on_planets[0], "N")
# Largest and smallest number in a list
print("The lightest a bus would be in the solar system is", bus_weight * min(gravity_on_planets), "N")
print("The heaviest a bus would be in the solar system is", bus_weight * max(gravity_on_planets), "N")
```
# Manipulate List Data
Python provides robust support for working with the data in lists. This support includes slicing data (examining just a portion) and sorting.
# Slicing List
```bash
planets = ["Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"]
planets_before_earth = planets[0:2]
print(planets_before_earth)
planets_after_earth = planets[3:]
# Joining List
amalthea_group = ["Metis", "Adrastea", "Amalthea", "Thebe"]
galilean_moons = ["Io", "Europa", "Ganymede", "Callisto"]
regular_satellite_moons = amalthea_group + galilean_moons
print("The regular satellite moons of Jupiter are", regular_satellite_moons)
# Sort list
regular_satellite_moons.sort() # Python sorts a list of strings in alphabetical order and a list of numbers in numeric order

```

While and For Loop In Python ::

- Identify when to use while and for loops.
- Run a task multiple times by using while loops.
- Loop over list data by using for loops.

Not : Python has many types that can be looped over. These types are known as iterables.
```bash
# While loop

# Create the variable for user input
user_input = ''
# Create the list to store the values
inputs = []

# The while loop
while user_input.lower() != 'done':
    # Check if there's a value in user_input
    if user_input:
        # Store the value in the list
        inputs.append(user_input)
    # Prompt for a new value
    user_input = input('Enter a new value, or done when done')
```

For Loop :

```bash
countdown = [4, 3, 2, 1, 0]
for number in countdown:
    print(number)
print("Blast off!! 🚀")
```
Manage data with Python dictionaries :

Python dictionaries allow you to model more complex data. Dictionaries are a collection of key/value pairs, and are very common in Python programs. Their flexibility allows you to dynamically work with related values without having to create classes or objects.

- Identify when to use a dictionary.
- Create and modify data inside a dictionary.
- Use dictionary methods to access dictionary data.

```bash
# Dictionary In Python
planet = {
    'name': 'Earth',
    'moons': 1
}

print(planet.get('name'))
wibble = planet.get('wibble') # Returns None !!
wibble = planet['wibble'] # Throws KeyError !!

#Although the behavior of get and the square brackets ([ ]) is generally the same for retrieving items, there's one key #difference. If a key isn't available, get returns None, and [ ] raises a KeyError.

planet.update({'name': 'Makemake'})
planet['name'] = 'Makemake'  # Syntactic sugar

# Multiple update at a time

planet.update({
    'name': 'Jupiter',
    'moons': 79
})

# Remove an item

planet.pop('orbital period')

# Complext data type :
# Dictionaries are able to store any type of a value, including other dictionaries. 

planet['diameter (km)'] = {
    'polar': 133709,
    'equatorial': 142984
}
print(f'{planet["name"]} polar diameter: {planet["diameter (km)"]["polar"]}')

```

# Dynamic programming with dictionary
```bash
# Retreive All Key Values
rainfall = {
    'october': 3.5,
    'november': 4.2,
    'december': 2.1
}

for key in rainfall.keys():
    print(f'{key}: {rainfall[key]}cm')

# Determine if a key exist in a dictionary
if 'december' in rainfall:
    rainfall['december'] = rainfall['december'] + 1
else:
    rainfall['december'] = 1

# Retreive All Keys
total_rainfall = 0
for value in rainfall.values():
    total_rainfall = total_rainfall + value
```

Functions ::

- Functions can require arguments or make them optional.
- You can extract reusable code and reuse it in a function.
- Variable arguments and variable keyword arguments are useful when you don't know the exact inputs.


```bash
# Create function
def rocket_parts():
    print("payload, propellant, structure")
# Function with arguments
def distance_from_earth(destination):
    if destination == "Moon":
        return "238,855"
    else:
        return "Unable to compute to that destination"

# Keyword Argument In Python
rom datetime import timedelta, datetime

def arrival_time(hours=51):
    now = datetime.now()
    arrival = now + timedelta(hours=hours)
    return arrival.strftime("Arrival: %A %H:%M")

# Mixing arguments & Keywork arguments
from datetime import timedelta, datetime

def arrival_time(destination, hours=51):
    now = datetime.now()
    arrival = now + timedelta(hours=hours)
    return arrival.strftime(f"{destination} Arrival: %A %H:%M")

# Use Variable Argument In Python
def variable_length(*args):
    print(args)

# Variable keyword arguments
def variable_length(**kwargs):
    print(kwargs)
```

Error Handling in Python ::

- Read and use error output from exceptions.
- Properly handle exceptions.
- Raise exceptions with useful error messages.
- Use exceptions to control a program's flow.

```bash
# try and catch
try:
     open('config.txt')
except FileNotFoundError:
     print("Couldn't find the config.txt file!")

# specific to generalized exception handler
def main():
    try:
        configuration = open('config.txt')
    except FileNotFoundError:
        print("Couldn't find the config.txt file!")
    except IsADirectoryError:
        print("Found config.txt but it is a directory, couldn't read it")
    except (BlockingIOError, TimeoutError):   # Grouping multiple exception together
        print("Filesystem under heavy load, can't complete reading configuration file")

# Accessing details about an exception
try:
    open("mars.jpg")
except FileNotFoundError as err:
     print("Got a problem trying to read the file:", err)

# access attributes of the error directly. 
try:
    open("config.txt")
except OSError as err:
     if err.errno == 2:
         print("Couldn't find the config.txt file!")
     elif err.errno == 13:
        print("Found config.txt but couldn't read it")

```








References ::
[link](https://learn.microsoft.com/en-us/training/paths/beginner-python/)

Check your knowledge ::  

1. What is the main version of Python currently supported? 
Python 3
2. What Python function is used to display information on the screen? 
print
3. What file extension is typically used for Python files? 
py
4. To show a string on the console, what code would you enter? 
The code print("my message")
5. If you invoke a program with program.py 1 2, what does the code sys.argv[0] contain? 
It contains program.py, the name of the program.
6. What will happen if you run the statement "1" + 2? 
It will give an error.
7. How do you know you're in a virtual environment named env? 
Your console prompt has changed to start with (env) ->.
8. How would you get a list of installed libraries? 
Run pip freeze.
9. How would you upgrade a package named "Flask"?
Run pip install Flask --upgrade.
10. Under what conditions does a test expression that uses the and operator evaluate to True? 
A test expression that uses the and operator evaluates to True when both subexpressions are true.
11. What is the keyword elif short for in Python? 
else if
12. What values in a test expression are always interpreted as false?
None and 0
# String
1. Are strings immutable?
Yes, because you can't modify them.
2. What types of quotation marks can be used with strings? 
Single, double, and triple quotation marks.
3. What is the string method call to use to break strings into separate lines?
The .split(' ') method.
4. What two string methods can you use to check to see whether a word exists in a string?
The .find() and .count() methods.
5. What is a problem you might run into when you use % for formatting?
The code can be harder to read when many replacements are needed.
6. How can you round down a number when you use f-strings?
Use the round() function inside the string, enclosed in braces ({}).
# Use Mathematical Operation In Python
1. What function is used to convert a string to an integer?
int
2. What function is used to determine an absolute value?
abs
3. What operator is used to determine a modulo?
%
# List Manipulation
1. What symbols are used to start and finish a list?
Brackets: []
2. Which statement would return the first value in the list planets?
planets[0]
3. Which statement would return the last value in the list planets?
planets[-1]
4. Which statement would sort the list planets? 
planets.sort()
# Loop
1. What statement would loop over the sequence collection and assign the variable item?
for item in collection
2. What is the correct statement to use to unpack variables named left and right from the list turns?
for left, right in turns
# Dictionary
1. What method allows access all key names in a Python dictionary?
keys
2. What method can you use to retrieve a value from a dictionary by using the key?
get
3. What method can you use to remove a key from a dictionary?
pop
# Function
1. Can a function that defines arguments be called without arguments?
No, because defined arguments are required.
2. Can a function with only keyword arguments be called without any arguments?
Yes, because keyword arguments are optional when calling a function.
3. Can a function define arguments and keyword arguments? 
Yes, but only if arguments are defined before keyword arguments.
4. What is the minimum number of arguments that a function can accept when you're using variable arguments?
Zero. There's no need to pass any argument at all when you're using variable arguments.
5. What is the syntax to declare variable arguments and variable keyword arguments? 
*args, **kwargs
# Exception
1. What are three elements that you might find in a traceback?
A file path, line number, and exception
2. What two keywords would you use for handling exceptions?
try and except
3. Why would using except Exception be unhelpful?
Because it can hide what the real problem is
4. When can it be useful to use as err in an except block?
When you want to reuse or inspect an exception
5. What is the right syntax to catch two exceptions in the same except line? 
except (ValueError, TypeError):






