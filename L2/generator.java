
import java.util.Random;
class generator 
{
	public static void main(String[] args)
	{
		Random random = new Random();
		System.out.println(1000);
		for(int i=0;i<1000;i++)
		{
			System.out.println(random.nextInt() % 10000);
		}
	}

}
