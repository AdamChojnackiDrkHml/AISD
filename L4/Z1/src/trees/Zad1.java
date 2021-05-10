package trees;

public class Zad1 {
    public static void main(String[] args) throws Exception {
        RBT bst = new RBT();

    //    bst.load("C:\\Users\\Adam\\Desktop\\AiSD\\L2\\AISD\\L4\\lab4\\aspell_wordlist.txt");
        bst.insert("aa");

        bst.insert("bab");
        bst.insert("cxb");
        bst.insert("ab");
        bst.insert("bbbbb");
       bst.insert("xdds");
       bst.insert("aaa");
       bst.insert("bbbtybb");
       bst.insert("bbuyjbb");
       bst.insert("bbnutbb");
//       bst.delete("bbb");
        bst.inorder();
//        bst.max(bst.Root);
    }
}
