package trees;

import java.io.File; 
import java.io.FileNotFoundException;
import java.util.Scanner; 

public class Splay
{


    public long comp;
    public class Node
    {
        public Node LeftSon;
        public Node RightSon;
        public int count;
        public String Value;

        public Node(String Value)
        {
            this.Value = Value;
            count = 1;
        }
    }

    
    Node Root;

    public void insert(String Value) 
    {
		Node N;
		int c;
		if (Root == null) 
        {
		    Root = new Node(Value);
		    return;
		}
		splay(Value);
		if ((c = Value.compareTo(Root.Value)) == 0) 
        {
		    //	    throw new DuplicateItemException(x.toString());
			Root.Value=Value;
            Root.count ++;
		    return;
		}
		N = new Node(Value);
		if (c < 0) 
        {
		    N.LeftSon = Root.LeftSon;
		    N.RightSon = Root;
		    Root.LeftSon = null;
		} 
        else 
        {
		    N.RightSon = Root.RightSon;
		    N.LeftSon = Root;
		    Root.RightSon = null;
		}
		Root = N;
		
	}


    private Node header = new Node(null);
    private void splay(String key) {
		Node l, r, t, y;
		l = r = header;
		t = Root;
		header.LeftSon = header.RightSon = null;
		for (;;) 
        {
            comp++;
		    if (key.compareTo(t.Value) < 0) 
            {
                if (t.LeftSon == null) 
                {
                    break;
                }
                comp++;
                if (key.compareTo(t.LeftSon.Value) < 0) 
                {
                    y = t.LeftSon;                            /* rotate right */
                    t.LeftSon = y.RightSon;
                    y.RightSon = t;
                    t = y;
                    if (t.LeftSon == null) break;
                }
                r.LeftSon = t;                                 /* link right */
                r = t;
                t = t.LeftSon;
		    } 
            else if (key.compareTo(t.Value) > 0) 
            {
                if (t.RightSon == null) 
                {
                    break;
                }
                comp++;
                if (key.compareTo(t.RightSon.Value) > 0) 
                {
                    y = t.RightSon;                            /* rotate left */
                    t.RightSon = y.LeftSon;
                    y.LeftSon = t;
                    t = y;
                    if (t.RightSon == null)
                    {
                        break;
                    }
                }
                l.RightSon = t;                                /* link left */
                l = t;
                t = t.RightSon;
		    } 
            else 
            {
                comp++;
		        break;
		    }
		}
		l.RightSon = t.LeftSon;                                   /* assemble */
		r.LeftSon = t.RightSon;
		t.LeftSon = header.RightSon;
		t.RightSon = header.LeftSon;
		Root = t;
	}
 

    Node find(String Value) 
    { 
        comp = 0;
        if (Root == null) 
        {
            return null;
        }

        splay(Value);
	    
        if(Root.Value.compareTo(Value) != 0) 
        {
            return null;
        }
        
        return Root;
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


    public void delete(String Value) 
    {
		Node x;
		splay(Value);
		if (Value.compareTo(Root.Value) != 0) 
        {
		    return;
		}
		// Now delete the root
		if (Root.LeftSon == null) 
        {
		    Root = Root.RightSon;
		} 
        else 
        {
		    x = Root.RightSon;
		    Root = Root.LeftSon;
		    splay(Value);
		    Root.RightSon = x;
		}

	}


    public Node max(Node N) 
    {
        Node x = N;
        if(N == null)
        {
            System.out.println();
            return null;
        }
        while(x.RightSon != null) 
        {
            x = x.RightSon;
        }
        splay(x.Value);
        System.out.println(x.Value);
        return x;		//Modified
    }


    public Node min(Node N) 
    {
        Node x = N;
        if(N == null) 
        {
            System.out.println();
            return null;
        }
        while(x.LeftSon != null) 
        {
            x = x.LeftSon;
        }
        splay(x.Value);
        System.out.println(x.Value);
        return x; //Modified
    }

    
    public Node successor(String value) 
    {
        Node K = find(value);
        Node X = K;
        if(K == null) 
        {
            System.out.println();
            return null;
        }

        if(K.RightSon != null) 
        {
            X = min(K.RightSon);
        }
        splay(X.Value);
        return X;
    }


    public void inorder()
    {
        inorder(Root);
    }
    public void inorder(Node N)
    {
        if(N == null)
        {
            System.out.println();
            return;
        }
        
        inorder(N.LeftSon);
        System.out.print(N.Value + " ");
        inorder(N.RightSon);

    }
    
}
