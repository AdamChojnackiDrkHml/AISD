package trees;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class RBT {

    public enum Color {
        RED,
        BLACK;
    }


    public class Node
    {
        public Node LeftSon;
        public Node RightSon;
        public Node Parent;
        public Color color; 
        public int count;
        private String Value;
        public boolean isNull;

        public Node(String Value)
        {
            this.Value = Value;
            count = 1;
            color = Color.RED;
        }

        public Node()
        {
            isNull = true;
            color = Color.BLACK;
            LeftSon = this;
            RightSon = this;
        }
    }



    
    Node Root;

    Node NullNode = new Node();
    

    public void insert(String NewValue) 
    {
        if(Root == null)
        {
            Root = new Node(NewValue);
            Root.color = Color.BLACK;
            createBond(Root, NullNode, true);
            createBond(Root, NullNode, false);
            return;
        }

        Node Current = Root;
        Node Parent = Current;
        boolean isLeft = false;
        while(!Current.isNull)
        {
            Parent = Current;
            if(Current.Value.compareTo(NewValue) == 0)
            {
                Current.count++;
                return;
            }

            if(Current.Value.compareTo(NewValue) < 0)
            {
                Current = Current.RightSon;
                if(Current.isNull)
                {
                    isLeft = false;
                }
            }
            else 
            {
                Current = Current.LeftSon;
                if(Current.isNull)
                {
                    isLeft = true;
                }
            }
        }
        Current = new Node(NewValue);


        createBond(Current, NullNode, true);
        createBond(Current, NullNode, false);

        createBond(Parent, Current, isLeft);


        
        repair(Current);
     
    }

    public void repair(Node N)
    {
        if(null == N)
        {
            return;
        }

        if(N.equals(Root) )
        {
            N.color = Color.BLACK;
            return;
        }

        if(!isRed(N.Parent))
        {
            return;
        }


        if(isRed(N.Parent) && isRed(brother(N.Parent)))
        {
            N.Parent.color = Color.BLACK;
            brother(N.Parent).color = Color.BLACK;
            N.Parent.Parent.color = Color.RED;
            
        }

        if(isRed(N.Parent) && !isRed(brother(N.Parent)))
        {
            if(N.Parent.equals(N.Parent.Parent.LeftSon)) 
            {
                if(N.equals(N.Parent.LeftSon))
                {
                    RotateLL(N);
                }
                else
                {
                    RotateRL(N);
                }
            }
            else
            {
                if(N.equals(N.Parent.LeftSon)) 
                {
                    RotateLR(N);
                } 
                else
                {
                    RotateRR(N);
                }
            }
        }
        repair(N.Parent.Parent);

    }

    private void RotateL(Node N)
    {
        if(!N.equals(Root)) {
            if(isLeftSon(N))
            {
                createBond(N.Parent, N.RightSon, true);
                createBond(N, N.RightSon.LeftSon, false);
                createBond(N.Parent.LeftSon, N, true);

            }
            else 
            {
                createBond(N.Parent, N.RightSon, false);
                createBond(N, N.RightSon.LeftSon, false);
                createBond(N.Parent.RightSon, N, true);
            }
        }
        else {

            Root = N.RightSon;
            Root.Parent = null;

            createBond(N, Root.LeftSon, false);
            createBond(Root, N, true);


        }
    }
    private void RotateR(Node N)
    {   
        if(!N.equals(Root)) {
            if(isLeftSon(N))
            {
                createBond(N.Parent, N.LeftSon, true);
                createBond(N, N.LeftSon.RightSon, true);
                createBond(N.Parent.LeftSon, N, false);

            }
            else 
            {
                createBond(N.Parent, N.LeftSon, false);
                createBond(N, N.LeftSon.RightSon, true);
                createBond(N.Parent.RightSon, N, false);
            }
        }
        else {
            Root = N.LeftSon;
            Root.Parent = null;

            createBond(N, Root.RightSon, true);
            createBond(Root, N, false);

        }
    }
    private void RotateLL(Node N)
    {
        RotateR(N.Parent.Parent);
        N.Parent.color = Color.BLACK;
        brother(N).color = Color.RED;
    }
    private void RotateRL(Node N)
    {
        RotateR(N.Parent);
        RotateRR(N);
    }
    private void RotateLR(Node N)
    {
        RotateL(N.Parent);
        RotateLL(N);
    }
    private void RotateRR(Node N)
    {
        RotateL(N.Parent.Parent);
        N.Parent.color = Color.BLACK;
        brother(N).color = Color.RED;
    }

    public boolean isRed(Node N)
    {
        return N.color.equals(Color.RED);
    }
    
    
    
    public Node brother(Node N)
    {

       if(N.equals(N.Parent.LeftSon))
       {
           return N.Parent.RightSon;
       }
       return N.Parent.LeftSon;

    }
    
    public boolean isLeftSon(Node N)
    {
        return N.equals(N.Parent.LeftSon);
    }
    public void load(String uri) 
    {
        try 
        {
            File myObj = new File(uri);
            Scanner myReader = new Scanner(myObj);
            while (myReader.hasNextLine()) 
            {
                String data = myReader.nextLine();
                char[] data1 = data.toCharArray();
                for (int i = 0; i < data1.length; i++){
                    if(!Character.isLetter(data1[i]))
                    {
                        data1[i] = ' ';
                    }
                }
                data = String.valueOf(data1);
                String[] words = data.split(" ");

                for(String word : words)
                {
                    if(!word.equals(""))
                    {
                        insert(word);
                    }
                }
            }
            myReader.close();
          } 
        catch (FileNotFoundException e) 
        {
            System.out.println("An error occurred.");
            e.printStackTrace();
        }
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

        if(K.LeftSon == null && K.RightSon != null)
        {
            if(K == Root)
            {
                Root = K.RightSon;
                return;
            }
            
            if(K == K.Parent.LeftSon)
            {
                K.Parent.LeftSon = K.RightSon;
                return;
            }
            
            if(K == K.Parent.RightSon)
            {
                K.Parent.RightSon = K.RightSon;
                return;
            }
        }

        if (K.LeftSon != null && K.RightSon == null)
        {
            if(K == Root)
            {
                Root = K.LeftSon;
                return;
            }
            
            if(K == K.Parent.LeftSon)
            {
                K.Parent.LeftSon = K.LeftSon;
                return;
            }
            
            if(K == K.Parent.RightSon)
            {
                K.Parent.RightSon = K.LeftSon;
                return;
            }
        }
        
        if(K.LeftSon != null && K.RightSon != null)
        {
            Node suc = successor(value);

            delete(suc.Value);

            K.Value = suc.Value;
            K.count = suc.count;
        }
    }


    public Node find(String value) 
    {
        if(Root == null)
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
            return min(K.RightSon);
        }

        Node Current = K.Parent;

        while ((Current != null) && K == Current.RightSon)
        {
            K = Current;
            Current = Current.Parent;
        }
        
        return Current;
    }

    public void inorder()
    {
        inorder(Root);
    }
    public void inorder(Node N)
    {
        if(N.isNull)
        {
 //           System.out.println();
            return;
        }
        System.out.print("[");
        inorder(N.LeftSon);
        System.out.print(N.Value + "(" + N.color +")");
        inorder(N.RightSon);
        System.out.print("]");

    }


    private void createBond(Node P, Node C, boolean isLeft)
    {
        if(isLeft)
        {
            P.LeftSon = C;
        }
        else
        {
            P.RightSon = C;
        }
        C.Parent = P;
    }
}
