
import java.util.Random;
class generator 
{
	public static void main(String[] args)
	{
		Random random = new Random();
		try 
		{
			for(int i=0;i<21;i++)
			{
				float dupa = random.nextFloat();
				while (dupa < 0.81)
				{
					dupa = random.nextFloat();
				}
				System.out.printf(String.format("%.2f", dupa) + "\n");
			}
		}
		catch (NumberFormatException ignore) {}

	}

}
