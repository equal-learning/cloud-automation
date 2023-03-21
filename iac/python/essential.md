# Python Essential - Two minute drills


Why python (what motivated creation of python)?

Its design philosophy emphasizes code readability, and its syntax allows programmers to express concepts
in fewer lines of code than would be possible in languages such as C++ or Java.

Coding philosophy:
Every named symbol that is declared has a type and a scope. Operations on the named symbol are allowed depending upon its type.

Python keywords (predefined named symbol) :

## Keywords  

    False class finally is  return None continue for lambda 
    try True def from nonlocal while and del global not
    with as elif if or yield assert else import pass break except in raise

## Operators  
    + - *  
    ** / //  
    %  
    << >>  
    & | ^  
    ~  
    < > <=  
    >= ==  
    !=  

## Python code structure:

* A python program is constructed from code blocks. A block isapiece of Python program text that isexecuted as a unit. The following are blocks: a module, a function body, and a class definition.

* Python code is structured as [script/function]/module/package

* Module: A module is a file containing Python definitions and statements. The file name is the module name with the suffix.pyappended. A module can contain executable statements as well as function definitions. These statements areintended to initialize the module. They are executed only thefirsttime the module name is encountered in animport statement. (They are also run if the file is executed as a script.)

* Executing a module as a script: By adding the following code at theend of the file, we can make the file usable as a script as well as an importablemodule.if __name__ == "__main__":
* Module search path: When a module namedspamis imported, the interpreter first searches for a built-in module with that name. If notfound, it then searches for a file namedspam.pyin a list of directories given by the variablesys.path. Thesys.pathis initialized from these locations:-The directory containing the input script (or the current directory).-PYTHONPATH (a list of directorynames, with the same syntax as the shell variable PATH).-the installation-dependent default.After initialization, Python programs can modify sys.path.

* Standard modules: Python comes with a library of standard modules.

* Package: Packages are a way ofstructuring Python’s module namespace by using “dotted module names”.A__init__.pyfile is required in the directory such that python can treat it as a package. In the simplest case,__init__.py can just be an empty file, but it can also execute initialization code for the package or set the__all__variable.

* Import statement: Import statement is used to import an entire package, module or selected class,functionor variable of a module. When a module is first imported, Python searches for the module andif found, it creates a module object,initializing it and binds to a name in the local scope. If the named module cannot be found, anImportErrorisraised.

Various import scenarios:

import foo # foo imported and bound locally
importfoo.bar.baz # foo.bar.baz imported, foo bound locally
import foo.bar.baz as fbb # foo.bar.baz imported and bound as fbb from foo.bar 
import baz # foo.bar.baz imported and bound as baz from foo 
import attr # foo importedand foo.attr bound as attr

* import *:
All public names defined in the module are bound in the local namespace for the scope where the import statementoccurs. Thepublic namesdefined by a module are determined by checking the module’s namespace for a variablenamed__all__; if defined, it must be a sequence of strings which are names defined or imported by that module.The names given in__all__are all considered public and are required to exist. If__all__is not defined, the set ofpublic names includesall names found in the module’s namespace which do not begin with an underscorecharacter ('_').__all__should contain the entire public API. It is intended to avoid accidentally exporting itemsthat are not part of the API (such as library modules whichwere imported and used within the module).

## Scope and Namespace:

Name spaceAnamespaceis a mapping fromnames to objects. Most namespaces are currently implemented as PythonDictionaries.The namespace for a module is automatically createdthe first time a module is imported.
Naming and binding:Names refer to object. Names are introduced by name binding operations.

Scope:
local
nonlocal
global

Resolution of names:A scope defines the visibility of a name within a block. If a local variable is defined in a block, its scopeincludes that block. If the definition occurs in a function block, the scope extends to any blocks containedwithin the defining one, unless a containing block introduces a different binding for the name.When a name isused in a code block, it is resolved using the nearest enclosing scope. The set of all suchscopes visible to a code block is called the block'environment.If theglobalstatement occurs within a block, all uses of the name specified in the statements refer to thebinding of that name in the top-level namespace. Names are resolved in the top-level namespace bysearching the global namespace, i,e. the namespace of the module containing the code block, and thebuiltins namespace.Thenonlocalstatement causes corresponding names to refer to previously bound variables in the nearestenclosing function scope.

## Function:

Call by value or call by reference : neither (call by object)

Lambda function: Smallanonymous functionscan be created with the lamda keyword. Lambda functionscan be used wherever function objects are required. They are syntactically restricted to a single expression.Like nested function definitions, lambda functions can reference variables from the containing scope.Ex-defsquare_root(x): returnmath.sqrt(x)using lambda function it can be written as:square_root = lambda x: math.sqrt(x)

## Classes:

Salient features of python classes (inspired from C++ and Module 3):

- Classes are created at runtime, andcan be modified further after creation!
- Class members arepublic(except for private variables).
- All member functions are virtual (in other words dynamic binding)
- Unlike java/C++; nothisreference available to access the members from the class methods. On the otherhand for each methodthe first argument(conventionally named asself)is used to represent a reference tothe current class instance (object) that can be used to access the members of the class.
- A class definition even can be placed inside the branch of anif statement or inside a method.
- When a class definition is entered,a new namespace is created and used as the local scope. Thus allassignments to local variables go into this new namespace. In particular, function definitions bind the nameof the new function here.When a class definition is leftnormally (via the end),aclass objectis created.This is basically a wrapper around the contents of the namespace created by the class definition.

Inheritenace:

- Supports multiple inheritance:class DerivedClassName(Base1, Base2, Base3)The search for attributes inherited from a parent class as depth-first, left-to-right, notsearching twice in the same class where there is an overlap in the hierarchy. Thus, if an attribute isnot found inDerivedClassName, it is searched for inBase1, then (recursively)in the base classesofBase1, and if it was not found there, it was searched for inBase2, and so on.
- An overriding in method in the derived class can extend rather simply replace the base method ofthe same name. To do so the following syntax isused :BaseClassName.methodname(self,arguments)

- There are two built in utility method that are useful with inheritance:
isinstance(object,classInfo) : To check an instance’s type.
issubclass(object,classInfo) : To check class inheritance.

Class Objects ( similar to java class object )

Class objects support two kinds of operations:

Attribute references and instantiation.
Ex
-
Consider the class
–
class MyClass:
"""A simple exa
mple class"""
i = 12345
def f(self):
return 'hello world'

Attribute reference operation
–
MyClass.i returns a interger.
MyClass.f returns a function object.

A function object is similar to static function in java and hence not bound to any particular instance of the class.

Instantiation Operation
–

x= MyClass ()
A new instance of MyClass object is created and assigned to x.
The instantiation operation (“calling” a class object) creates an empty object. Therefore a class
may define a special method named __init__() to provide a particular initial state to the class object. Unlike
java and c++, there is no way to declare constructor in python !
Ex
–
def __init__(self):
self.data = []

* Instance Object
The only operations understood by instance objects are attribute references. There are two kinds of validattribute names--data attributes and methods.Data attribute:Ex-x.i refers to the member integer i.Method attribute:x.f referes to the method object f.

* Method Object 

For every function object (ex-MyClass.f)there is on method object which is associated to a particularinstance of the class.Method objects are used without specifying the self attribute.Method objects can access the global variables.

* Private variables
“Private” instance variables that cannot be accessed except from inside an objectdon’t exist in Python.However, there isa convention that is followed by most Python code: a name prefixed with an underscore(e.g._spam) should be treated as a non-public part of the API (whether it is a function, a method or a datamember). It should be considered an implementation detail andsubject to change without notice.

Iterators:
The use of iterators for the container objects (Iterables) pervades and unifies Python.

Ex
–
    for element in [1, 2, 3]:
    print(element):

Here behind the scene the for statement calls iter() on the container object. The function returns an iterator object
that defines the method __next__() which accesses elements in the container one at a time. When there are no more
elements, __next__() raises a StopIteration exception which tells the for loop to terminate. You can call the
__next__() method using the next() built
-
in function.
Iterator can easily be added for a class by defining an __iter__() method which returns an object with a __next__()
method. If the class defines __next__(), then __iter__() can just return self.

class Reverse:
"""Iterator for looping over a sequence backwards."""
def __init__(self, data):
self.data = data
self.index = len(data)
def __iter__(self):
return self
def __next__(self):
if self.index == 0:
raise StopIteration
self.index = self.index - 1
return self.data[self.index]

Generators (also referred as coroutine):
Generator functions allow to createa function that behaves likes an iterable; in other words it can generate a seriesof values.Generators are a simple and powerful toolfor creating iterators. They are written like regular functions but usetheyieldstatement whenever they want to return data (note that returningdata is not same as returning the contrl.).Each timenext()is called on it, the generator resumes where it left-off (it remembers all the data values and whichstatement was last executed). An example shows that generators can be trivially easy to create:

def reverse(data):
for index in range(len(data)-1,-1,-1):
yield data[index]
for char in reverse('golf'):
print(char)

Discussion: [Some interesting discussions on generator & coroutine](https://www.jeffknupp.com/blog/2013/04/07/improve-your-python-yield-and-generators-explained/)

Essential buit-in [function](https://docs.python.org/3/library/functions.html)

Standard Library :

os : For interacting with the operating system.

Shutil : For daily file and directory management tasks.

Glob: For file listing of a directory using wildcard searches.

Sys:
For processingcommand linearguments.
For outputredirection.
Programtermination

re:Provides regularexpression tools for stringprocessing.

math: Providesaccess to math functions.

random: Provides tool for random selection.

urllib.request : Provides tools for internet access.

datetime: Supplies classes for date and time manipulation.
zlib, gzip, bz2, lzma,zipfile and tarfile: For data compression and archiving.
timeit: Tools for measuring code performance.
Threading : For threads.


Ref:

Python call by value or call by reference : neither (call by object)
[link](https://jeffknupp.com/blog/2012/11/13/is-python-callbyvalue-or-callbyreference-neither/)

Python execution model:
[link](https://docs.python.org/3.4/reference/executionmodel.html)

Python FAQ:
[link](https://docs.python.org/3.4/faq/)




