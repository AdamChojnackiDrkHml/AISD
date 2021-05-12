package trees;

import java.io.File; 
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner; 

public class BST{


    public long comp;

    public class Node
    {
        public Node LeftSon;
        public Node RightSon;
        public Node Parent;
        public int count;
        private String Value;

        public Node(String Value)
        {
            this.Value = Value;
            count = 1;
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
        
        if(isLeft)
        {
            Current.Parent = Parent;
            Parent.LeftSon = Current;
            return;
        }
        Current.Parent = Parent;
        Parent.RightSon = Current;

    }


    public List<String> load(String uri) 
    {
        List<String> list = new ArrayList<String>();
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
                        list.add(word);
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
        return list;
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
        comp = 0;
        if(Root == null)
        {
            System.err.println(0);
            return null;
        }

        Node Current = Root;
        boolean isFound = false;
        while(Current != null)
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
        System.out.println(Current.Value);
        return Current;
    }

    public void inorder()
    {
        if(Root != null)
        {
            inorder(Root);
            return;
        }
        System.out.println();
    }
    public void inorder(Node N)
    {
        if(N == null)
        {
            System.out.print("");
            return;
        }
        inorder(N.LeftSon);
        System.out.print(N.Value + " ");
        inorder(N.RightSon);

    }
    
}
