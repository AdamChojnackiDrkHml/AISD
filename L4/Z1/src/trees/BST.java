package trees;

public class BST{



    private class Node
    {
        public Node LeftSon;
        public Node RightSon;
        public Node Parent;

        private String Value;
    
        public Node(String Value)
        {
            this.Value = Value;
        }
    }
    
    
    Node Root;


    public void insert(String NewValue) 
    {
        if(Root == null)
        {
            Root = new Node(NewValue);
            return;
        }

        Node Current = Root;
        Node Parent = Current;
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
        Current = new Node(NewValue);
        
        if(isLeft)
        {
            Parent.LeftSon = Current;
            return;
        }

        Parent.RightSon = Current;

    }


    public void load(String uri) 
    {
        // TODO Auto-generated method stub
    }


    public void delete(String value)
    {
        Node K = find(value);

        if(K.LeftSon == null && K.RightSon == null)
        {
            if(K == K.Parent.LeftSon)
            {
                K.Parent.LeftSon = null;
                return;
            }
            
            if(K == K.Parent.RightSon)
            {
                K.Parent.RightSon = null;
            }
        }

        if((K.LeftSon == null && K.RightSon != null )
        ||(K.LeftSon != null && K.RightSon == null)
        {
            if(K == Root)
            {
            
            }
        }
         
    }


    public Node find(String value) 
    {
        if(Root != null)
        {
            System.err.println();
            return null;
        }

        Node Current = Root;
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
        return Current;

    }


    public Node min(Node N) 
    {
        if(N == null)
        {
            System.out.println();
            return null;
        }

        Node Min = N;

        while(Min.LeftSon != null)
        {
            Min = Min.LeftSon;
        }

        System.out.println(Min.Value);
        return Min;

    }


    public Node max(Node N) 
    {
        if(N == null)
        {
            System.out.println();
            return null;
        }

        Node Max = N;

        while(Max.RightSon != null)
        {
            Max = Max.RightSon;
        }

        System.out.println(Max.Value);
        return Max;
    }


    public Node successor(String value) 
    {
        Node K = find(value);
        if(K == null) 
        {
            System.out.println();
            return null;
        }

        if(K.RightSon != null) 
        {
            return max(K.RightSon);
        }

        Node Current = K.Parent;

        while ((Current != null) && K == Current.RightSon)
        {
            K = Current;
            Current = Current.Parent;
        }
        
        return Current;
    }


    public void inorder(Node N)
    {
        if(N == null)
        {
            System.out.println();
            return;
        }

        inorder(N.LeftSon);
        System.out.println(N.Value);
        inorder(N.RightSon);

    }
    
}
