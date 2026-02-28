// https://leetcode.com/problems/design-hashset/

// Using a vector of list (applying the concept of chaining in hashing to avoid collisions)
class MyHashSet {
public:
    vector<list<int>> store;
    int size;
    MyHashSet() {
        size=100;
        store.resize(size);
    }
    
    int hash(int key){
        return key%size;
    }
    
    list<int> :: iterator search(int key){
        int i=hash(key);
        return find(store[i].begin(),store[i].end(),key);
    }
    
    void add(int key) {
        if(contains(key))   return;
        int i=hash(key);
        store[i].push_back(key);
    }
    
    void remove(int key) {
        if(!contains(key))   return;
        int i=hash(key);
        store[i].erase(search(key));
    }
    
    bool contains(int key) {
        int i=hash(key);
        if(search(key)!=store[i].end()) return true;
        else return false;
    }
};