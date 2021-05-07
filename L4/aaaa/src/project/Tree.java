package Project;

public interface Tree 
{
    boolean insert(String value);
    boolean load(String uri);
    boolean delete(String value);
    boolean find(String value);
    boolean min();
    boolean max();
    boolean successor(String value);
    boolean inorder();
}
