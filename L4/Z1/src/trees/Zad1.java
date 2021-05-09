package trees;

public class Zad1 {
    public static void main(String[] args) throws Exception {
        BST bst = new BST();

    //    bst.load("C:\\Users\\Adam\\Desktop\\AiSD\\L2\\AISD\\L4\\lab4\\aspell_wordlist.txt");
        bst.insert("aa");

        bst.insert("bbb");
        bst.insert("bbb");
        bst.insert("ab");
        bst.insert("bbbbb");
       bst.insert("xdds");
       bst.insert("aa");
       bst.insert("bbbbb");
       bst.insert("bbbbb");
       bst.insert("bbbbb");
//       bst.delete("bbb");
        bst.inorder();
//        bst.max(bst.Root);
    }
}
