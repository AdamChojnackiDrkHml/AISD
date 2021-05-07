package project;


public interface TreeI 
{
    void insert(String value);
    void load(String uri);
    void delete(String value);
    void find(String value);
    void min();
    void max();
    void successor(String value);
    void inorder();
}
