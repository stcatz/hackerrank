t = int(raw_input())

for ts in range(t):
	m, n = [int(i) for i in raw_input().strip().split()]
	hc = [int(i) for i in raw_input().strip().split()]
	vc = [int(i) for i in raw_input().strip().split()]

	hc.sort()
	hc.reverse()

	vc.sort()
	vc.reverse()

	hfactor, vfactor, result = 1, 1, 0

	while len(hc)!=0 or len(vc)!=0:
		if len(hc) == 0:
			#print vc, vfactor
			result += sum(map(lambda x: x*vfactor, vc))
			result %= 1000000007
			break
		elif len(vc) == 0:
			#print hc, hfactor
			result += sum(map(lambda x: x*hfactor, hc))
			result %= 1000000007
			break
		
		#inc_h = sum(map(lambda x: x*(vfactor+1), vc))

		#inc_v = sum(map(lambda x: x*(hfactor+1), hc))

		if vc[0] > hc[0]:
			#print vc[0], vfactor
			result += vc[0] * vfactor
			result %= 1000000007
			vc.remove(vc[0])
			hfactor += 1
		else:
			#print hc[0], hfactor
			result += hc[0] * hfactor
			result %= 1000000007
			hc.remove(hc[0])
			vfactor += 1
		#print vc, vfactor, inc_v, " = ", hc, hfactor, inc_h


	print result
