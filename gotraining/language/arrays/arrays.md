## Arrays

Arrays are a special data structure in Go that allow us to allocate contiguous blocks of fixed size memory. Arrays have some special features in Go related to how they are declared and viewed as types.
数组是 Go 中的一种特殊数据结构，它允许我们分配固定大小的连续内存块。数组在 Go 中有一些与如何声明和查看类型相关的特殊功能。
## Notes

* If you don't understand the data, you don't understand the problem.
* 如果你不理解数据，你就不会理解问题。
* 
* If you don't understand the cost of solving the problem, you can't reason about the problem.
* 如果您不了解解决问题的成本，就无法对问题进行推理。
* 
* If you don't understand the hardware, you can't reason about the cost of solving the problem.
* 如果您不了解硬件，则无法推理解决问题的成本。
* 
* Arrays are fixed length data structures that can't change.
* 数组是固定长度的数据结构，不能改变。
* 
* Arrays of different sizes are considered to be of different types.
* 不同大小的数组被认为是不同的类型。
* 
* Memory is allocated as a contiguous block.
* 内存被分配为一个连续的块。
* 
* Go gives you control over spacial locality.
* Go 使您可以控制空间局部性。


## Code Review 代码练习

[Declare, initialize and iterate 声明、初始化和迭代](example/declare-initialize/declare-initialize.go)  
[Different type arrays](example/different-type/different-type.go)  
[Contiguous memory allocations](example/contiguous-memory-allocations/contiguous-memory-allocations.go)    
[Range mechanics](example/range/range.go)  

## Exercises

### Exercise 1

Declare an array of 5 strings with each element initialized to its zero value. Declare a second array of 5 strings and initialize this array with literal string values. Assign the second array to the first and display the results of the first array. Display the string value and address of each element.
声明一个包含 5 个字符串的数组，每个元素都初始化为零值。  
声明第二个包含 5 个字符串的数组，并使用文字字符串值初始化该数组。  
将第二个数组分配给第一个数组并显示第一个数组的结果。  
显示每个元素的字符串值和地址。  

[Exercise](exercise/exercise.go)


## CPU Cache/Memory相关资料   

## CPU Caches 待学习

[CPU Caches and Why You Care (18:50-20:30)](https://youtu.be/WDIkqP4JbkE?t=1129) - Scott Meyers  
[CPU Caches and Why You Care (44:36-45:40)](https://youtu.be/WDIkqP4JbkE?t=2676) - Scott Meyers   
[Performance Through Cache-Friendliness (4:25-5:48)](https://youtu.be/jEG4Qyo_4Bc?t=266) - Damian Gryski

## CPU Cache Notes

* CPU caches works by caching main memory on cache lines.
* Cache lines today are either 32 or 64 bytes wide depending on the hardware.
* Cores do not access main memory directly. They tend to only have access their local caches.
* Both data and instructions are stored in the caches.
* Cache lines are shuffled down L1->L2->L3 as new cache lines need to be stored in the caches.
* Hardware likes to traverse data and instructions linearly along cache lines.
* Main memory is built on relatively fast cheap memory. Caches are built on very fast expensive memory.

* Access to main memory is incredibly slow, we need the cache.
  * Accessing one byte from main memory will cause an entire cache line to be read and cached.
  * Writes to one byte in a cache line requires the entire cache line to be written.

* Small = Fast
  * Compact, well localized code that fits in cache is fastest.
  * Compact data structures that fit in cache are fastest.
  * Traversals touching only cached data is the fastest.

* Predictable access patterns matter.
  * Whenever it is practical, you want to employ a linear array traversal.
  * Provide regular patterns of memory access.
  * Hardware can make better predictions about required memory.

* Cache misses can result in TLB cache misses as well.
  * Cache of translations of a virtual address to a physical address.
  * Waiting on the OS to tell us where the memory is.

#### CPU Caches / Memory 待学习

[CPU Caches and Why You Care - Video](https://www.youtube.com/watch?v=WDIkqP4JbkE) - Scott Meyers  
[A Crash Course in Modern Hardware - Video](https://www.youtube.com/watch?v=OFgxAFdxYAQ) - Cliff Click  
[NUMA Deep Dive Series](http://frankdenneman.nl/2016/07/06/introduction-2016-numa-deep-dive-series/) - Frank Denneman

[CPU Caches and Why You Care - Deck](http://www.aristeia.com/TalkNotes/codedive-CPUCachesHandouts.pdf) - Scott Meyers  
[Mythbusting Modern Hardware to Gain 'Mechanical Sympathy'](https://www.youtube.com/watch?v=MC1EKLQ2Wmg) - Martin Thompson  
[What Every Programmer Should Know About Memory](http://www.akkadia.org/drepper/cpumemory.pdf) - Ulrich Drepper  
[How CPU Caches Work and Why](http://www.extremetech.com/extreme/188776-how-l1-and-l2-cpu-caches-work-and-why-theyre-an-essential-part-of-modern-chips) - Joel Hruska  
[Modern Microprocessors A 90 Minute Guide](http://www.lighterra.com/papers/modernmicroprocessors) - Jason Robert Carey Patterson  
[Memory part 2: CPU caches](http://lwn.net/Articles/252125) - Ulrich Drepper  
[The Free Lunch Is Over](http://www.gotw.ca/publications/concurrency-ddj.htm) - Herb Sutter  
[Data Center Computers: Modern Challenges in CPU Design](https://m.youtube.com/watch?feature=youtu.be&v=QBu2Ae8-8LM) - Dick Sites  
[Wirth's Law](https://en.wikipedia.org/wiki/Wirth%27s_law) - Wikipedia  
[Eliminate False Sharing](http://www.drdobbs.com/parallel/eliminate-false-sharing/217500206) - Herb Sutter  
[The Myth Of Ram](http://www.ilikebigbits.com/2014_04_21_myth_of_ram_1.html) - Emil Ernerfeldt  
[Understanding Transaction Hardware Memory](https://www.infoq.com/presentations/hardware-transactional-memory) - Gil Gene  
[Performance Through Cache-Friendliness (4:25-5:48)](https://youtu.be/jEG4Qyo_4Bc?t=266) - Damian Gryski  
[Going Nowhere Faster](https://www.youtube.com/watch?v=2EWejmkKlxs) - Chandler Carruth

#### Data-Oriented Design 待学习

[Data-Oriented Design and C++](https://www.youtube.com/watch?v=rX0ItVEVjHc) - Mike Acton  
[Efficiency with Algorithms, Performance with Data Structures](https://www.youtube.com/watch?v=fHNmRkzxHWs) - Chandler Carruth  
[Taming the performance Beast](https://www.youtube.com/watch?v=LrVi9LHP8Bk) - Klaus Iglberger

[Pitfalls of OOP](http://harmful.cat-v.org/software/OO_programming/_pdf/Pitfalls_of_Object_Oriented_Programming_GCAP_09.pdf) - Tony Albrecht  
[Why you should avoid Linked Lists](https://www.youtube.com/watch?v=YQs6IC-vgmo) - Bjarne Stroustrup  
[Data-Oriented Design (Or Why You Might Be Shooting Yourself in The Foot With OOP)](http://gamesfromwithin.com/data-oriented-design) - Noel    
[Was object-oriented programming a failure?](https://www.quora.com/Was-object-oriented-programming-a-failure) - Quora
