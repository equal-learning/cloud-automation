
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
