
public class Prime {

	public static boolean isPrime(long num) {
		if (num<2) return false;
		long i = 2;
		while(i*i<=num) {
			if (num%i==0) return false;
			i += 2;
			if (i<5) i--;
		}
		return true;
	}

	public static void main(String[] args) {
		for(int i=1;i<1000;i++) {
			if (isPrime(i)) {
				System.out.println("prime: " + i);
			}
		}	
	}
}
