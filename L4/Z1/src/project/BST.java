package project;

public class BST implements TreeI{

    private BST Parent;
    private BST LeftSon;
    private BST RightSon;

    private String Value;

    public BST(String Value) {
        this.Value = Value;
    }

    @Override
    public boolean insert(String NewValue) 
    {
        if(NewValue.compareTo(this.Value) == 0) 
        {
            
        }
        return false;
    }

    @Override
    public boolean load(String uri) 
    {
        // TODO Auto-generated method stub
        return false;
    }

    @Override
    public boolean delete(String value)
     {
        // TODO Auto-generated method stub
        return false;
    }

    @Override
    public boolean find(String value) 
    {
        // TODO Auto-generated method stub
        return false;
    }

    @Override
    public boolean min() 
    {
        // TODO Auto-generated method stub
        return false;
    }

    @Override
    public boolean max() 
    {
        // TODO Auto-generated method stub
        return false;
    }

    @Override
    public boolean successor(String value) 
    {
        // TODO Auto-generated method stub
        return false;
    }

    @Override
    public boolean inorder()
    {
        // TODO Auto-generated method stub
        return false;
    }
    
}
