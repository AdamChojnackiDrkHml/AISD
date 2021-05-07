package project;

public class BST implements TreeI{



    private class BSTNode 
    {
        private BSTNode Parent;
        private BSTNode LeftSon;
        private BSTNode RightSon;
    
        private String Value;
    
        public BSTNode(String Value)
        {
            this.Value = Value;
        }
    }
    
    
    BSTNode Root;

    @Override
    public void insert(String NewValue) 
    {
        if(Root == null)
        {
            Root = new BSTNode(NewValue);
            return;
        }

        BSTNode Current = Root;
        BSTNode Parent = Current;
        boolean isLeft = false;
        while(Current != null)
        {
            Parent = Current;

            if(Current.Value.compareTo(NewValue) < 0)
            {
                Current = Current.RightSon;
                
                if(Current == null)
                {
                    isLeft = true;
                }

            }
            else 
            {
                Current = Current.LeftSon;
            }
        }
        Current = new BSTNode(NewValue);
        
        if(isLeft)
        {
            Parent.LeftSon = Current;
            return;
        }

        Parent.RightSon = Current;

    }

    @Override
    public void load(String uri) 
    {
        // TODO Auto-generated method stub
    }

    @Override
    public void delete(String value)
     {
        // TODO Auto-generated method stub
         
    }

    @Override
    public void find(String value) 
    {
        if(Root != null)
        {
            System.err.println("Nonexisting tree");
            return;
        }

        BSTNode Current = Root;
        boolean isFound = false;
        while(Current != null)
        {
            if(Current.Value.compareTo(value) < 0)
            {
                Current = Current.RightSon;
            }
            else if(Current.Value.compareTo(value) > 0)
            {
                Current = Current.LeftSon;
            }
            else 
            {
                isFound = true;
                break;
            }
        }

        System.out.println(isFound ? 1 : 0);

    }

    @Override
    public void min() 
    {
        if(Root == null)
        {
            System.out.println();
            return;
        }

        BSTNode Min = Root;

        while(Min.LeftSon != null)
        {
            Min = Min.LeftSon;
        }

        System.out.println(Min.Value);

    }

    @Override
    public void max() 
    {
        // TODO Auto-generated method stub

    }

    @Override
    public void successor(String value) 
    {
        // TODO Auto-generated method stub

    }

    @Override
    public void inorder()
    {
        // TODO Auto-generated method stub

    }
    
}
