## C++ OOP Notes (with Examples)

This document summarizes C++ OOP concepts and points you to concrete examples in this folder.

Compile examples with:

```bash
g++ -std=c++20 -Wall -Wextra -pedantic -O2 <file>.cpp -o <file>
./<file>
```

---

### 1. Classes and Objects

- **Concept**: A *class* is a user-defined type that groups data (members) and functions (methods). An *object* is an instance of a class.
- **Key points**:
  - Syntax: `class Name { /* members */ };`
  - Objects can be created on the **stack** or **heap**.
  - Member functions operate on the data of the object.
- **Example file**: `01_classes_objects.cpp`

---

### 2. Encapsulation & Access Specifiers

- **Concept**: Encapsulation hides internal data and exposes only safe operations (interface).
- **Access specifiers**:
  - `public`: accessible from anywhere.
  - `private`: accessible only inside the class (and friends).
  - `protected`: accessible in the class and derived classes.
- **Typical pattern**:
  - Private data members.
  - Public member functions (getters/setters or behavior).
- **Example file**: `02_encapsulation_access.cpp`

---

### 3. Constructors, Destructors, and `this`

- **Constructors**:
  - Default, parameterized, copy, move, delegating constructors.
  - Use **member initializer lists** to initialize data.
- **Destructor**:
  - `~ClassName()` called when object lifetime ends (stack scope ends or `delete` on heap).
  - Used to release resources (memory, files, locks).
- **`this` pointer**:
  - Inside non-static member functions, `this` points to the current object.
  - Use `this->member` when names conflict or for chaining (`return *this;`).
- **Example file**: `03_ctors_dtors_this.cpp`

---

### 4. Static Members (Class-Level Data/Functions)

- **Static data members**:
  - Shared across all objects of the class.
  - Declared inside class, defined once outside.
- **Static member functions**:
  - Do not have `this`.
  - Can be called as `ClassName::func()`.
- **Use cases**:
  - Object counters, configuration values, helpers that don't depend on a specific object.
- **Example file**: `04_static_members.cpp`

---

### 5. Friend Functions and Friend Classes

- **Concept**: `friend` gives a non-member function or another class access to private/protected members.
- **Use cases**:
  - Symmetric operators (like `operator<<` for streams).
  - Helpers that need deep access but should not be member functions.
- **Caution**: Overuse breaks encapsulation; use sparingly.
- **Example file**: `05_friend_functions_classes.cpp`

---

### 6. Operator Overloading

- **Concept**: Give intuitive meanings to operators for user-defined types.
- **Rules**:
  - At least one operand must be of a user-defined type.
  - Do not abuse: keep semantics natural (e.g., `+` means addition-like behavior).
- **Common operators**:
  - Arithmetic: `+`, `-`, `*`, `/`, `%`
  - Comparison: `==`, `!=`, `<`, `>`, `<=`, `>=`
  - Assignment family: `=`, `+=`, `-=`, etc.
  - I/O: `<<`, `>>`
- **Example file**: `06_operator_overloading.cpp`

---

### 7. Inheritance

- **Concept**: Build new classes (derived) from existing ones (base) to reuse and extend behavior.
- **Types of inheritance**:
  - `public` inheritance: models an **is-a** relationship.
  - `protected` / `private` inheritance: used for implementation details.
- **Guideline**: For typical OOP hierarchies, use `public` inheritance.
- **Example file**: `07_inheritance.cpp`

---

### 8. Polymorphism (Static and Dynamic)

- **Static (compile-time) polymorphism**:
  - Achieved via function overloading and templates.
  - Resolved at compile time.
- **Dynamic (run-time) polymorphism**:
  - Achieved via virtual functions and base-class pointers/references.
  - Resolved at runtime using virtual tables (vtables).
- **Example file**: `08_polymorphism.cpp`

---

### 9. Virtual Functions, `override`, `final`, and Abstract Classes

- **Virtual function**:
  - Declared with `virtual` in base class.
  - Overridden in derived classes.
- **`override` keyword**:
  - Tells the compiler you intend to override a virtual function; helps catch mistakes.
- **`final` keyword**:
  - Prevents further overriding or inheriting from a class.
- **Abstract class**:
  - Has at least one **pure virtual** function (`= 0`).
  - Cannot be instantiated; used as an interface or base type.
- **Example file**: `09_virtual_abstract.cpp`

---

### 10. Multiple Inheritance and Virtual Inheritance

- **Multiple inheritance**:
  - A class inherits from more than one base class.
  - Useful for combining interfaces or behaviors.
- **Diamond problem**:
  - Ambiguity when a class inherits from two classes that share a common base.
  - Solved using **virtual inheritance**.
- **Syntax**:
  - `class Derived : public Base1, public Base2 { ... };`
  - `class B : virtual public A { ... };`
- **Example file**: `10_multiple_inheritance.cpp`

---

### 11. Object Lifetime, Copy Semantics, and Move Semantics

- **Object lifetime**:
  - Stack objects: lifetime is the scope.
  - Heap objects: lifetime controlled by `new/delete` (or smart pointers).
- **Copy semantics**:
  - Copy constructor: `ClassName(const ClassName&)`.
  - Copy assignment: `ClassName& operator=(const ClassName&);`
- **Move semantics** (C++11+):
  - Move constructor: `ClassName(ClassName&&)`.
  - Move assignment: `ClassName& operator=(ClassName&&);`
  - Used to efficiently transfer ownership of resources instead of copying.
- **Rule of 3/5**:
  - If you implement one of destructor, copy ctor, copy assignment, you often need all three.
  - With move semantics, this extends to five (add move ctor and move assignment).
- **Example file**: `11_copy_move_semantics.cpp`

---

### 12. RAII (Resource Acquisition Is Initialization)

- **Concept**:
  - Acquire resources (memory, file handles, locks) in constructor.
  - Release them in destructor.
  - Ensures no leaks, even with exceptions or early returns.
- **Examples in standard library**:
  - `std::unique_ptr`, `std::shared_ptr`
  - `std::lock_guard`, `std::scoped_lock`
- **Example file**: `12_raii.cpp`

---

### 13. Smart Pointers and Polymorphism

- **`std::unique_ptr`**:
  - Exclusive ownership.
  - Cannot be copied, only moved.
- **`std::shared_ptr`**:
  - Shared ownership via reference counting.
  - Use when multiple owners are required.
- **`std::weak_ptr`**:
  - Non-owning reference to a `shared_ptr` resource (breaks cycles).
- **With polymorphism**:
  - Store base-class pointers: `std::unique_ptr<Base>`.
  - Enable dynamic dispatch without manual `new/delete`.
- **Example file**: `13_smart_pointers_polymorphism.cpp`

---

### 14. OOP + Templates (Advanced)

- **Templates** give compile-time polymorphism.
- Combine with OOP when:
  - You want generic containers of polymorphic objects.
  - You want zero-overhead abstractions for performance.
- **Example file**: `14_templates_and_oop.cpp`

---

### 15. pImpl (Pointer-to-Implementation) Idiom (Advanced Encapsulation)

- **Goal**:
  - Hide implementation details in `.cpp` files.
  - Reduce compile times and maintain binary compatibility.
- **Pattern**:
  - Public class holds a `std::unique_ptr<Impl>` where `Impl` is defined in the `.cpp`.
- **Example file**: `15_pimpl_idiom.cpp` (header-style demo inside one file for learning)

---

### 16. RTTI and Casting (`dynamic_cast`, `static_cast`, `reinterpret_cast`, `const_cast`)

- **RTTI (Run-Time Type Information)**:
  - `dynamic_cast` and `typeid` work correctly only with **polymorphic** types (at least one virtual function).
- **`dynamic_cast`**:
  - Safe downcast; returns `nullptr` for invalid pointer cast, throws `std::bad_cast` for reference cast.
- **`static_cast`**:
  - Compile-time cast; no runtime checking, faster but can be unsafe if misused.
- **`reinterpret_cast`**:
  - Low-level, implementation-defined reinterpretation of bits; avoid unless you know exactly what you’re doing.
- **`const_cast`**:
  - Adds/removes `const`/`volatile` qualifiers; never use to modify truly-constant objects.
- **Example file**: `16_rtti_casting.cpp`

---

### 17. Virtual Destructors and Object Slicing

- **Virtual destructor**:
  - Any class intended to be used polymorphically (deleted through a base pointer) should have a **virtual destructor**.
  - Ensures derived destructor runs and resources are released correctly.
- **Object slicing**:
  - Happens when a derived object is **copied into** a base object (not reference/pointer).
  - Derived parts are “sliced off”; virtual dispatch then only sees base subobject.
- **Guidelines**:
  - Use references/pointers (`Base&`, `Base*`, smart pointers) when working with polymorphic types.
  - Avoid passing/returning polymorphic objects by value.
- **Example file**: `17_virtual_destructor_slicing.cpp`

---

### 18. Defaulted and Deleted Special Member Functions

- **`= default`**:
  - Ask the compiler to generate the default implementation (even if you also define others).
- **`= delete`**:
  - Explicitly disable a function (e.g., non-copyable classes).
- **Use cases**:
  - Non-copyable types (e.g., `std::unique_ptr`-like).
  - Enforce construction rules (e.g., delete default constructor, allow only parameterized).
- **Example file**: `18_default_delete_special_members.cpp`

---

### 19. Exception Safety and RAII

- **Exception safety guarantees**:
  - *No-throw*: operation will not throw.
  - *Strong*: on failure, state is unchanged (commit/rollback).
  - *Basic*: invariants preserved; state may change but remains valid.
- **RAII and exceptions**:
  - Use RAII to automatically clean up on exceptions (destructors always run).
  - Prefer strong exception safety when writing modifying functions (copy-and-swap idiom).
- **Example file**: `19_exception_safety_raii.cpp`

---

### 20. Composition vs Inheritance and SOLID-ish Design

- **Composition over inheritance**:
  - Prefer “has-a” (member object) over “is-a” (public inheritance) when reusing implementation.
- **When to use inheritance**:
  - To model a true **is-a** relationship with a stable base interface (`Shape`, `Device`, etc.).
- **SOLID-flavored ideas**:
  - Keep interfaces small and focused.
  - Open for extension, closed for modification (prefer new derived types over editing base code).
  - Depend on abstractions (interfaces) rather than concrete classes.
- **Example file**: `20_composition_vs_inheritance.cpp`

---

### Suggested Study Order

1. `01_classes_objects.cpp`
2. `02_encapsulation_access.cpp`
3. `03_ctors_dtors_this.cpp`
4. `04_static_members.cpp`
5. `05_friend_functions_classes.cpp`
6. `06_operator_overloading.cpp`
7. `07_inheritance.cpp`
8. `08_polymorphism.cpp`
9. `09_virtual_abstract.cpp`
10. `10_multiple_inheritance.cpp`
11. `11_copy_move_semantics.cpp`
12. `12_raii.cpp`
13. `13_smart_pointers_polymorphism.cpp`
14. `14_templates_and_oop.cpp`
15. `15_pimpl_idiom.cpp`
16. `16_rtti_casting.cpp`
17. `17_virtual_destructor_slicing.cpp`
18. `18_default_delete_special_members.cpp`
19. `19_exception_safety_raii.cpp`
20. `20_composition_vs_inheritance.cpp`

For each file:

- Read the code.
- Run it and modify values / add print statements.
- Try to write a small similar example yourself without looking.

