package trees;

import java.util.Random;
import java.util.Scanner;


public class Zad1 {

    static enum treeType {
        BST,
        RBT,
        SPLAY;
    }
    public static void main(String[] args) throws Exception {

        BST bst = new BST();
        RBT rbt = new RBT();
        Splay splay = new Splay();
        treeType tree = treeType.BST;
        Scanner scanner = new Scanner(System.in);
        int n = 0;
        try {
            n = Integer.parseInt(scanner.next());
        } catch (NumberFormatException e)
        {
            return;
        }
        switch(args[1])
        {
            case "bst": 
            {
                tree = treeType.BST;
                break;
            }
            case "rbt":
            {   
                tree = treeType.RBT;
                break;
            }
            case "splay": 
            {
                tree = treeType.SPLAY;
                break;
            }

        }
        switch (tree) 
        {
            case BST:
                while(n > 0)
                {
                    String command = scanner.next();

                    switch (command)
                    {
                        case "max":
                        {
                            bst.max(bst.Root);
                            break;
                        }
                        case "min":
                        {
                            bst.min(bst.Root);
                            break;
                        }
                        case "succesor":
                        {
                            bst.successor(scanner.next());
                            break;
                        }
                        case "inorder":
                        {
                            bst.inorder();
                            System.out.println();
                            break;
                        }
                        case "find":
                        {
                            bst.find(scanner.next());
                            break;
                        }
                        case "insert":
                        {
                            bst.insert(scanner.next());
                            break;
                        }
                        case "load":
                        {
                            bst.load(scanner.next());
                            break;
                        }
                        case "delete":
                        {
                            bst.delete(scanner.next());
                            break;
                        }
                        default:
                        {
                            break;
                        }

                    }
                    n--;
                }
                break;
            case RBT:
            while(n != 0)
            {
                String command = scanner.next();

                switch (command)
                {
                    case "max":
                    {
                        rbt.max(rbt.Root);
                        break;
                    }
                    case "min":
                    {
                        rbt.min(rbt.Root);
                        break;
                    }
                    case "succesor":
                    {
                        rbt.successor(scanner.next());
                        break;
                    }
                    case "inorder":
                    {
                        rbt.inorder();
                        System.out.println();
                        break;
                    }
                    case "find":
                    {
                        rbt.find(scanner.next());
                        break;
                    }
                    case "insert":
                    {
                        rbt.insert(scanner.next());
                        break;
                    }
                    case "load":
                    {
                        rbt.load(scanner.next());
                        break;
                    }
                    case "delete":
                    {
                        rbt.delete(scanner.next());
                        break;
                    }
                    default:
                    {
                        break;
                    }

                }
            }
            break;
            case SPLAY:
            while(n != 0)
            {
                String command = scanner.next();

                switch (command)
                {
                    case "max":
                    {
                        splay.max(splay.Root);
                        break;
                    }
                    case "min":
                    {
                        splay.min(splay.Root);
                        break;
                    }
                    case "succesor":
                    {
                        splay.successor(scanner.next());
                        break;
                    }
                    case "inorder":
                    {
                        splay.inorder();
                        System.out.println();
                        break;
                    }
                    case "find":
                    {
                        splay.find(scanner.next());
                        break;
                    }
                    case "insert":
                    {
                        splay.insert(scanner.next());
                        break;
                    }
                    case "load":
                    {
                        splay.load(scanner.next());
                        break;
                    }
                    case "delete":
                    {
                        splay.delete(scanner.next());
                        break;
                    }
                    default:
                    {
                        break;
                    }

                }
            }
            break;
            default:
                break;

        }
        System.out.println();
    }


}
