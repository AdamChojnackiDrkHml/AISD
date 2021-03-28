
import java.util.Random;
class generator 
{
	public static void main(String[] args)
	{
		Random random = new Random();
		try 
		{
			int n = Integer.parseInt(args[0]);
			System.out.println(n);

			for(int i=0;i<n;i++)
			{
				System.out.println(random.nextInt() % n);
			}
		}
		catch (NumberFormatException ignore) {}

	}

}
