package trees;

public class Zad1 {
    public static void main(String[] args) throws Exception {
        RBT bst = new RBT();

        bst.load("/home/adam/Uczelnia/AiSD/L4/Z1/src/trees/lotr.txt");
        bst.loadDelete("/home/adam/Uczelnia/AiSD/L4/Z1/src/trees/lotr.txt");
        bst.inorder();
        System.out.println();
    }
}
