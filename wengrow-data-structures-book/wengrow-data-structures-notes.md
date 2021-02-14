# Wengrow: A Common-Sense Guide to Data STructures and Algorithms

## Ch. 1: Why Data Structures Matter
- arrays
- big four operations on a data structure: 
  - reading (constant time in array)
  - searching (N time in array)
  - insertion (constant time at end of array, N time at beginning or middle?)
  - deletion (N time in array-- worst case)
- sets
  - insertion requires a search.

## Ch. 2: Why Algorithms Matter
- some algorithms can be faster than others by orders of magnitude
- it is helpful to know how to count steps of an algorithm
- binary search can finish in O(log N) time. Linear search works in O(N) time
- binary search only works on ordered arrays
## Ch. 3: O Yes! Big O Notation
- O notation works in relation to N elements. The key question is always, "how will my algorithm perform as the input grows larger?", or "how will the algorithm's performance change as the data increases?"
- logarithms are key to O and speeding up algorithms. A logarithm is the measure of how many steps a number has to be divided in two (for log base 2) until we get to the number 1. What this means practically is that an algorithm that performs at O(log N) time is very fast, especially for large data sets. This is because as data sets become larger and larger, a doubling of the data set results in a much larger number N, while only decreasing the algorithm's performance by 1 step.

## Ch. 4: Speeding Up Your Code with Big O
- introduces O(N^2) vs. O(N)
- bubble sort: O(N^2)

## Ch. 5: Optimizing Code with and Without Big O
- O() notation ignores constants
- while two algorithms can have the same O() performance, one can still be faster than the other because O() ignores constants.
- selection sort -- faster than bubble sort, but both are O(N^2).

## Ch. 6: Optimizing for Optimistic Scenarios
- O notation is usually describing worst-case scenarios
- but we can also evaluate best-case, or average-case scenarios.
- Also, O notation only takes into account the highest-order of N if there are multiple orders present.
- insertion sort: O(N^2)
- Q: when doing O() calculations, why don't we take into account the implementation of the underlying array? From what I remember, some programming languages use linked lists instead of arrays -- wouldn't this affect the performance of something like insertion sort? I suppose though if we're always iterating over the whole array then the efficiency will be the same. It's only if we need to continuously get random addresses in the middle of a linked list that we will see a slowdown in performance...I think.
- worst-case vs. average scenario vs. best case for insertion sort: worst case is O(N^2), but average case is still the same in terms of O(), although it would be N^2 / 2. Best case is O(N). Selection sort is still O(N^2) for best case, but the same as insertion sort for the average case and a little better for worst case.
- the takeaway is that if you have an idea of how sorted your data will be, you might want to choose selection or insertion sort depending.

## Ch. 7: Big O in Everyday Code
- mean average
- word builder
- array sample
- average celsius reading
- clothing labels
- count the ones
- palindrome checker
- get all the products
- multiple datasets p. 105-6: two arrays of different lengths interacting through multiplication: we express it as `n * m`. For this example, it could be O(N^2) or O(N), somewhere in that range
- password cracker: an algorithm that computes every possible combination of alphabetic letters for a given length of string, N. As N increases, the number of possible combinations increases by 26^N. For example, if we are looking for all combinations of a 3-letter password, we need to multiply `26*26*26`. This is very slow: it's O(2^N), or in this case, O(26^N)


## Ch. 8: Blazing Fast Lookup with Hash Tables
- hash tables work by 'hashing' the keys, that is, converting the key into a more-or-less unique number using a hashing algorithm, then storing the value at that index in an array. 
- Because of this, lookup using a hash table occurs in constant time.
- But it's only one-directional. If we want to see if a value exists, or if we want to check if a key exists for a particular value, that's going to take O(N) time.
- Collisions, or different keys whose hashed values are the same, can happen with hash tables. In this case, the table will store subarrays with key/value pairs inside them. Lots of collisions slows down the speed of a hash table because searching for a key in subarrays takes O(N) time. To avoid too many collisions, the rule of thumb for the `load factor`, or ratio of data to cells, is .7 (10 cells for 7 pieces of data), which balances memory usage with collisions.
- Hash tables/dicts/maps (in golang) 

## Ch. 9: Crafting Elegant Code with Stacks and Queues
- these are constrained data structures. The question is why would we ever want this, because it seems to offer only a subset of the functionality of something like an array? WELL they can prevent potential bugs; they give us concrete mental models; 
- stacks are LIFO (last in first out). Insertion, deletion, and read can only occur at the top of the stack. The bottom of the stack is at the other end.
  - checking for matching braces.
- queues are FIFO (first in first out). Insertion can only occur at the end of the queue. Deletion and read can only occur at the front of the queue.

## Ch. 10: Recursively Recurse with Recursion

## Ch. 11: Learning to Write in Recursive

## Ch. 12: Dynamic Programming
- problem: when we have two (or more) recursive calls in a function, the function has an efficiency of 0(2^N) which is terrible. This is because for each additional data element, the number of function calls doubles and so the algorithmic steps double as well. 
- two approaches as a solution: memoization or bottom-up
- memoization saves the result of a calculation in a hash table.
- bottom up uses iteration instead of recursion.
- either approach can result in an efficiency of 0(N)
- DO THE EXERCISES

## Ch. 13: Recursive Algorithms for Speed
- partitioning: the process of taking a random element from an array and making sure that all values smaller are to the left, and that all values larger are to the right. Then that element is correctly sorted.
- quicksort: this is review for me -- his description is not exactly the same as what I remember, but it makes sense. The partition algorithm seems more complicated than the one I learned before. But here it is: establish pivot point; establish pointers at the left and the right sides of the array (excluding the pivot); move the left pointer to the right and compare to the pivot, and keep moving as long as the left pointer is less than the pivot. Then move the right pointer to the left and compare to the pivot, and keep moving as long as the right pointer is greater than the pivot; once both pointers have stopped moving, we see if they have reached the same point or crossed paths, and if they have, we swap the value at the left pointer with the pivot. If they haven't, we swap the values at left and right pointers and restart the process.
- quickselect: ?

## Ch. 14: Node-Based Data Structures
- a node is a data piece, that can live anywhere in memory
- linked list (p. 238 for comparison table)
  - reading (worst-case: N time)
  - searching (N time)
  - insertion (constant time at beginning of list -- much better than array)
  - deletion (constant time at beginning of list)
- doubly linked list
  - has immediate access to the end of the list
  - can traverse backwards
  - because they have access to front and end and can insert and remove in constant time to these two elements, they are good as queues.

## Ch. 15: Speeding Up All the Things With Binary Search Trees
- binary search trees solve a problem presented by hash tables and arrays: hash tables have fast search, insertion, and deletion, but are unordered. Arrays are ordered and has fast reads and search, but slow insertion and deletion (O(N)).
- binary search trees are ordered, with fast insertion, deletion, and search.
- binary search trees have `root` nodes, and include terms `parent`, `child`, `ancestor`, `descendant`.
- binary search trees have `levels`, which are the nodes within a row in the tree
- they are said to be `balanced` or `imbalanced` depending on if the subtrees have the same number of nodes in them.
- binary trees are trees in which each node has zero, one, or two child nodes.
- binary search trees fulfill the following condition: 
  - every node has at most a "left" and a "right" child.
  - all "left" descendants of a node must contain values that are less than the node
  - all "right" descendants of a node must contain values that are greater than the node
- when binary search trees are balanced (this holds true for binary trees in general), there will be logN levels in the tree. This is because each level doubles the number of nodes. 
- This also means that search is a O(log N) operation for binary search trees
- Insertion is basically the same operation as search, with the extra step of actually inserting the data. So it's O(log N) as well.
- This means that binary search trees are great for applications where there will be lots of changes to data.
- a completely imbalanced tree takes O(N) time for search and insertion. This occurs when already ordered data is converted into a tree and is the worse-case scenario. Ideally, data should be randomized when it is being inserted into a binary search tree. Then it takes O(log N) time.
- deletion is the most complex of the four primary operations (search, read (is this the same as search? In an array it's not), insert(write), delete)
  - if the node being deleted has no children, delete it
  - if the node being deleted has one child, replace the node with the child. The child is called the "successor" node.
  - if the node being deleted has two children, follow this rule:
    - replace the node with the least of all of the nodes with greater values.
      - in practice, this means we traverse downward from the node to be deleted: first, to the "right" child node, then to the left again and again until we reach the bottom node. This value is, by definition, going to satisfy the above condition and becomes the successor node.
      - however, the successor node might have a right child of its own. What do we do in this case?
      - In this case, we place the right child where the successor node used to be. In other words, we take the right child of the successor node and make it the left child of the successor node's old parent.
- in-order traversal:
  - recursively call `inOrderTraversal` on the left node; visit the current node; call `inOrderTraversal` on the right node; returns undefined/null
  - according to the above, the order of 'visiting' will be: leftmost node; the parent of the leftmost node; the right child of that followed by the leftmost descendent. This way, the nodes are traversed left-to-right, that is, in order of least to greatest. Reversing the order of this would traverse in greatest-to-least order.
  
## Ch. 16: Keeping Your Priorities Straight with Heaps
  - a priority queue is a queue in which insertion keeps the order of the elements. In a classic array, this means O(N) insertion (since elements need to be shifted to the right) along with O(1) deletion. So we use a heap to make that faster
  - a binary heap is a kind of binary tree. Within that there's a max-heap and a min-heap. We're talking about the binary max-heap here.
  - two conditions to a heap:
    - the value of a node must be greater than each of its descendant nodes. Called the `heap condition`. A min-heap has the opposite condition.
    - the tree must be complete. This means that every row must have all nodes except the bottom row. The bottom row can have missing nodes, but there must not be a node to the right of a missing node.
  - the heap is weakly ordered, which means search is not a fast operation, unlike with a binary search tree. 
  - BUT heaps are primarily used for the operations of insertion and deletion, which are what we want in a priority queue.
  - heaps have a `last node`, which is the rightmost node on the bottom level.
  - for insertion, we insert a new node as a new 'last node'. Then we compare it to its parent; if it's greater, we swap. We continue on this way until no swap is made or until it becomes the root element. This process of swapping is called `trickling up`. The time complexity is O(log N)
  - for deletion, we only ever delete the root. To do this, we replace the root with the last node. Then we `trickle down` the new root to its proper place. To trickle down:
    - check both children for the larger of the two;
    - if the trickle down node is smaller than the larger of the two, swap them.
    - repeat the two steps above until no swap is made.
  - the time complexity of insertion is O(log N)
  - to solve the problem of the last node (how do we find which node is last?), we can implement a heap AS an array(!). For me, the question is, why doesn't this create the same problem as in the beginning? If the root node is simply the first element of the array, then deletion is still O(N) operation, right?