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

        public Node(String Value)
        {
            this.Value = Value;
            count = 1;
            color = Color.RED;
        }
    }

    
    Node Root;


    public void insert(String NewValue) 
    {
        if(Root == null)
        {
            Root = new Node(NewValue);
            Root.color = Color.BLACK;
            return;
        }

        Node Current = Root;
        Node Parent = Current;
        boolean isLeft = false;
        while(Current != null)
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
            }
            else 
            {
                Current = Current.LeftSon;
                if(Current == null)
                {
                    isLeft = true;
                }
            }
        }
        Current = new Node(NewValue);
        Current.Parent = Parent;
        if(isLeft)
        {
            Parent.LeftSon = Current;
        }
        if(!isLeft)
        {
            Parent.RightSon = Current;
        }


        try {
            repair(Current);
        } catch (Exception e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }



    }

    public void repair(Node N) throws Exception
    {
        if(!isRed(N.Parent))
        {
            return;
        }

        if(brother(N) != null && isRed(brother(N)))
        {
            N.Parent.color = Color.BLACK;
            brother(N).color = Color.BLACK;

            if(N.Parent.Parent != null && N.Parent.Parent != Root)
            {
                N.Parent.Parent.color = Color.RED;
            }
            return;
        }

        if(N.Parent.Parent.RightSon == N.Parent && N.Parent.RightSon == N)
        {
            
        }

    }

    public boolean isRed(Node N) throws Exception
    {
        if(N != null)
        {
            return N.color == Color.RED;
        }
        throw new Exception("YOU VIOLATED THE LAW");
    }
    public Node brother(Node N)
    {
        if(N != null && N.Parent != null)
        {
            if(N == N.Parent.LeftSon && N.Parent.RightSon != null)
            {
                return N.Parent.RightSon;
            }
            if(N == N.Parent.RightSon && N.Parent.LeftSon != null)
            {
                return N.Parent.LeftSon;
            }
        }
        return null;
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
        if(N == null)
        {
 //           System.out.println();
            return;
        }
        System.out.print("[");
        inorder(N.LeftSon);
        System.out.print(N.Value + "(" + N.count +")");
        inorder(N.RightSon);
        System.out.print("]");

    }
}
