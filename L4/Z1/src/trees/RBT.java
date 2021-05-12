package trees;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class RBT {

    public enum Color {
        RED,
        BLACK;
    }

    public long comp;
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


        
        repairInsert(Current);
     
    }

    public void repairInsert(Node N)
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
                    RotateLR(N);
                }
            }
            else
            {
                if(N.equals(N.Parent.LeftSon)) 
                {
                    RotateRL(N);
                } 
                else
                {
                    RotateRR(N);
                }
            }
        }
        try {
        repairInsert(N.Parent.Parent);
        }
        catch (NullPointerException e)
        {
            return;
        }

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
        RotateRR(N.RightSon);
    }
    private void RotateLR(Node N)
    {
        RotateL(N.Parent);
        RotateLL(N.LeftSon);
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

    public void loadDelete(String uri) 
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

                            delete(word);
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

    private void rbTransplant(Node u, Node v) {
        if (u.Parent == null) {
          Root = v;
        } else if (u == u.Parent.LeftSon) {
          u.Parent.LeftSon = v;
        } else {
          u.Parent.RightSon = v;
        }
        v.Parent = u.Parent;
      }

    public void delete(String value)
    {
        
        Node K = find(value);
        if(K.isNull)
        {
            return;
        }
        Node X,Y;
        Y = K;
        Color originalColor = K.color;
        if(K.LeftSon.isNull)
        {
            X = K.RightSon;
            rbTransplant(K, K.RightSon);
        }
        else if(K.RightSon.isNull)
        {
            X = K.LeftSon;
            rbTransplant(K,K.LeftSon);
        }
        else
        {
            Y = min(K.RightSon);
            originalColor = Y.color;
            X = Y.RightSon;

            if(Y.Parent.equals(K))
            {
                X.Parent = Y;
            }
            else
            {
                rbTransplant(Y, Y.RightSon);
                createBond(Y, K.RightSon, false);
            }
            
            rbTransplant(K, Y);
            createBond(Y, K.LeftSon, true);
            Y.color = K.color;

            
        }
        
        if(originalColor.equals(Color.BLACK))
        {
            repairDelete(X);
        }
    }


    private void repairDelete(Node X)
    {
        Node s;
        while (X != Root && X.color.equals(Color.BLACK)) 
        {
          if (X.equals(X.Parent.LeftSon)) 
          {
            s = X.Parent.RightSon;
            if (s.color.equals(Color.RED)) 
            {
              s.color = Color.BLACK;
              X.Parent.color = Color.RED;
              RotateL(X.Parent);
              s = X.Parent.RightSon;
            }
    
            if (s.LeftSon.color.equals(Color.BLACK) && s.RightSon.color.equals(Color.BLACK)) 
            {
              s.color = Color.RED;
              X = X.Parent;
            } 
            else 
            {
              if (s.RightSon.color.equals(Color.BLACK)) 
              {
                s.LeftSon.color = Color.BLACK;
                s.color = Color.RED;
                RotateR(s);
                s = X.Parent.RightSon;
              }
    
              s.color = X.Parent.color;
              X.Parent.color = Color.BLACK;
              s.RightSon.color = Color.BLACK;
              RotateL(X.Parent);
              X = Root;
            }
          } 
          else 
          {
            s = X.Parent.LeftSon;
            if (s.color.equals(Color.RED)) 
            {
              s.color = Color.BLACK;
              X.Parent.color = Color.RED;
              RotateR(X.Parent);
              s = X.Parent.LeftSon;
            }
    
            if (s.RightSon.color.equals(Color.BLACK) && s.LeftSon.color.equals(Color.BLACK)) 
            {
              s.color = Color.RED;
              X = X.Parent;
            } 
            else 
            {
              if (s.LeftSon.color.equals(Color.BLACK)) 
              {
                s.RightSon.color = Color.BLACK;
                s.color = Color.RED;
                RotateL(s);
                s = X.Parent.LeftSon;
              }
    
              s.color = X.Parent.color;
              X.Parent.color = Color.BLACK;
              s.LeftSon.color = Color.BLACK;
              RotateR(X.Parent);
              X = Root;
            }
          }
        }
        X.color = Color.BLACK;
    }
    public Node find(String value) 
    {

        comp = 0;
        if(Root == null)
        {
            System.err.println(0);
            return null;
        }

        Node Current = Root;
        boolean isFound = false;
        while(!Current.isNull)
        {
            comp++;
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
        if(N.isNull)
        {
            System.out.println();
            return null;
        }

        Node Min = N;

        while(!Min.LeftSon.isNull)
        {
            Min = Min.LeftSon;
        }

        System.out.println(Min.Value);
        return Min;

    }


    public Node max(Node N) 
    {
        if(N.isNull)
        {
            System.out.println();
            return null;
        }

        Node Max = N;

        while(!Max.RightSon.isNull)
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

        if(K.RightSon.isNull) 
        {
            return min(K.RightSon);
        }

        Node Current = K.Parent;
        
        while ((Current != null) && K == Current.RightSon)
        {
            K = Current;
            Current = Current.Parent;
        }
        System.out.println(Current.Value);
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
            System.out.println();
            return;
        }
        inorder(N.LeftSon);
        System.out.print(N.Value + "(" + N.color +") ");
        inorder(N.RightSon);

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
